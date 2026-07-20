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
	inspectionApp "prahari/services/inspection/internal/application/inspection"
	checklistApp "prahari/services/inspection/internal/application/checklist"
	actionApp "prahari/services/inspection/internal/application/action"
	scheduleApp "prahari/services/inspection/internal/application/schedule"
	searchApp "prahari/services/inspection/internal/application/search"
	reportingApp "prahari/services/inspection/internal/application/reporting"
	exportApp "prahari/services/inspection/internal/application/export"

	// Infrastructure adapters
	pgInfra "prahari/services/inspection/internal/infrastructure/postgres"
	redisInfra "prahari/services/inspection/internal/infrastructure/redis"
	kafkaInfra "prahari/services/inspection/internal/infrastructure/kafka"
	workflowInfra "prahari/services/inspection/internal/infrastructure/workflow"
	identityInfra "prahari/services/inspection/internal/infrastructure/identity"
	permitInfra "prahari/services/inspection/internal/infrastructure/permit"
	incidentInfra "prahari/services/inspection/internal/infrastructure/incident"
	eventsIface "prahari/services/inspection/internal/interfaces/events"

	// Interfaces
	httpIface "prahari/services/inspection/internal/interfaces/http"
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

	tp, err := InitTelemetry(ctx, "inspection-service", cfg.Environment)
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
	identityClient := identityInfra.NewClient(cfg.IdentityGrpcAddr)
	_ = permitInfra.NewClient(cfg.PermitGrpcAddr)
	incidentClient := incidentInfra.NewClient(cfg.IncidentGrpcAddr)

	inspectionStore := pgInfra.NewInspectionStore(db)
	templateStore := pgInfra.NewTemplateStore(db)
	checklistStore := pgInfra.NewChecklistStore(db)
	itemStore := pgInfra.NewItemStore(db)
	findingStore := pgInfra.NewFindingStore(db)
	observationStore := pgInfra.NewObservationStore(db)
	capaStore := pgInfra.NewCAPAStore(db)
	commentStore := pgInfra.NewCommentStore(db)
	attachmentStore := pgInfra.NewAttachmentStore(db)
	timelineStore := pgInfra.NewTimelineStore(db)
	auditStore := pgInfra.NewAuditStore(db)
	scheduleStore := pgInfra.NewScheduleStore(db)

	_ = redisInfra.NewCache(redisClient)
	kafkaPublisher := kafkaInfra.NewPublisher(nil)

	// DI Injection
	inspectionSvc := inspectionApp.NewService(inspectionStore, kafkaPublisher, auditStore, timelineStore, workflowClient, incidentClient)
	checklistSvc := checklistApp.NewService(templateStore, checklistStore, itemStore)
	actionSvc := actionApp.NewService(capaStore, auditStore, timelineStore)
	scheduleSvc := scheduleApp.NewService(scheduleStore)
	searchSvc := searchApp.NewService(inspectionStore)
	reportingSvc := reportingApp.NewService(nil)
	exportSvc := exportApp.NewService(inspectionStore)

	kafkaConsumer := kafkaInfra.NewConsumer(inspectionStore)
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	mux := http.NewServeMux()
	httpIface.RegisterRoutes(mux, inspectionSvc, checklistSvc, actionSvc, scheduleSvc, searchSvc, reportingSvc, exportSvc)

	pipelineHandler := InitRouter(mux)
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "Inspection Management Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down Inspection Management Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	eventListener.Stop(shutdownCtx)
	_ = srv.Shutdown(shutdownCtx)

	prahariLogger.Info(ctx, "Inspection Management Service exited gracefully.")
}
