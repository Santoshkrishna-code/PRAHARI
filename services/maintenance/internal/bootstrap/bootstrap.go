package bootstrap

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	prahariLogger "prahari/shared/logger"

	// Application services
	maintenanceApp "prahari/services/maintenance/internal/application/maintenance"
	workorderApp "prahari/services/maintenance/internal/application/workorder"
	schedulingApp "prahari/services/maintenance/internal/application/scheduling"
	planningApp "prahari/services/maintenance/internal/application/planning"
	searchApp "prahari/services/maintenance/internal/application/search"
	reportingApp "prahari/services/maintenance/internal/application/reporting"
	exportApp "prahari/services/maintenance/internal/application/export"

	// Infrastructure adapters
	pgInfra "prahari/services/maintenance/internal/infrastructure/postgres"
	redisInfra "prahari/services/maintenance/internal/infrastructure/redis"
	kafkaInfra "prahari/services/maintenance/internal/infrastructure/kafka"
	workflowInfra "prahari/services/maintenance/internal/infrastructure/workflow"
	identityInfra "prahari/services/maintenance/internal/infrastructure/identity"
	assetInfra "prahari/services/maintenance/internal/infrastructure/asset"
	permitInfra "prahari/services/maintenance/internal/infrastructure/permit"
	eventsIface "prahari/services/maintenance/internal/interfaces/events"

	// Interfaces
	httpIface "prahari/services/maintenance/internal/interfaces/http"
)

// RunApp handles bootstrap coordination sequence.
func RunApp() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := LoadConfig(ctx)
	if err != nil {
		panic(err)
	}

	_ = InitLogger(cfg.Environment)

	tp, err := InitTelemetry(ctx, "maintenance-service", cfg.Environment)
	if err == nil && tp != nil {
		defer func() { _ = tp.Shutdown(ctx) }()
	}

	db, err := InitDatabase(ctx, cfg.DatabaseURL)
	if err != nil {
		prahariLogger.Error(ctx, "Database initialization failed", prahariLogger.Err(err))
	}
	if db != nil {
		defer db.Close()
	}

	redisClient, err := InitRedis(ctx, cfg.RedisAddr)
	if err == nil && redisClient != nil {
		defer redisClient.Close()
	}

	_ = InitKafka(ctx, cfg.KafkaBrokers)

	workflowClient := workflowInfra.NewClient(cfg.WorkflowGrpcAddr)
	_ = identityInfra.NewClient(cfg.IdentityGrpcAddr)
	assetClient := assetInfra.NewClient(cfg.AssetGrpcAddr)
	permitClient := permitInfra.NewClient(cfg.PermitGrpcAddr)

	maintenanceStore := pgInfra.NewMaintenanceStore(db)
	workorderStore := pgInfra.NewWorkOrderStore(db)
	planStore := pgInfra.NewPlanStore(db)
	scheduleStore := pgInfra.NewScheduleStore(db)
	_ = pgInfra.NewTaskStore(db)
	_ = pgInfra.NewPartStore(db)
	_ = pgInfra.NewLaborStore(db)
	_ = pgInfra.NewDowntimeStore(db)
	_ = pgInfra.NewChecklistStore(db)
	_ = pgInfra.NewCommentStore(db)
	timelineStore := pgInfra.NewTimelineStore(db)
	auditStore := pgInfra.NewAuditStore(db)

	_ = redisInfra.NewCache(redisClient)
	kafkaPublisher := kafkaInfra.NewPublisher(nil)

	// DI Injection
	maintenanceSvc := maintenanceApp.NewService(maintenanceStore, kafkaPublisher, auditStore, timelineStore, assetClient, permitClient, workflowClient)
	workorderSvc := workorderApp.NewService(workorderStore)
	schedulingSvc := schedulingApp.NewService(scheduleStore)
	planningSvc := planningApp.NewService(planStore)
	searchSvc := searchApp.NewService(maintenanceStore)
	reportingSvc := reportingApp.NewService(nil)
	exportSvc := exportApp.NewService(maintenanceStore)

	kafkaConsumer := kafkaInfra.NewConsumer(maintenanceStore)
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	mux := http.NewServeMux()
	httpIface.RegisterRoutes(mux, maintenanceSvc, workorderSvc, schedulingSvc, planningSvc, searchSvc, reportingSvc, exportSvc)

	pipelineHandler := InitRouter(mux)
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "Maintenance Management Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down Maintenance Management Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	eventListener.Stop(shutdownCtx)
	_ = srv.Shutdown(shutdownCtx)

	prahariLogger.Info(ctx, "Maintenance Management Service exited gracefully.")
}
