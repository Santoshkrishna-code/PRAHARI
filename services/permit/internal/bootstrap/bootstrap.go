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
	permitApp "prahari/services/permit/internal/application/permit"
	approvalApp "prahari/services/permit/internal/application/approval"
	riskApp "prahari/services/permit/internal/application/riskassessment"
	searchApp "prahari/services/permit/internal/application/search"
	reportingApp "prahari/services/permit/internal/application/reporting"
	exportApp "prahari/services/permit/internal/application/export"

	// Infrastructure adapters
	pgInfra "prahari/services/permit/internal/infrastructure/postgres"
	redisInfra "prahari/services/permit/internal/infrastructure/redis"
	kafkaInfra "prahari/services/permit/internal/infrastructure/kafka"
	workflowInfra "prahari/services/permit/internal/infrastructure/workflow"
	identityInfra "prahari/services/permit/internal/infrastructure/identity"
	incidentInfra "prahari/services/permit/internal/infrastructure/incident"
	eventsIface "prahari/services/permit/internal/interfaces/events"

	// Interfaces
	httpIface "prahari/services/permit/internal/interfaces/http"
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

	tp, err := InitTelemetry(ctx, "permit-service", cfg.Environment)
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
	_ = incidentInfra.NewClient(cfg.IncidentGrpcAddr)

	permitStore := pgInfra.NewPermitStore(db)
	typeStore := pgInfra.NewTypeStore(db)
	approvalStore := pgInfra.NewApprovalStore(db)
	riskStore := pgInfra.NewRiskStore(db)
	hazardStore := pgInfra.NewHazardStore(db)
	isolationStore := pgInfra.NewIsolationStore(db)
	gasStore := pgInfra.NewGasStore(db)
	searchStore := pgInfra.NewSearchStore(db)
	auditStore := pgInfra.NewAuditStore(db)
	timelineStore := pgInfra.NewTimelineStore(db)

	_ = redisInfra.NewCache(redisClient)
	kafkaPublisher := kafkaInfra.NewPublisher(nil)

	// DI Injection
	permitSvc := permitApp.NewService(permitStore, kafkaPublisher, auditStore, timelineStore, workflowClient)
	approvalSvc := approvalApp.NewService(approvalStore, identityClient, kafkaPublisher, auditStore, timelineStore, permitStore)
	riskSvc := riskApp.NewService(riskStore, hazardStore, isolationStore, gasStore, auditStore, timelineStore)
	searchSvc := searchApp.NewService(searchStore)
	reportingSvc := reportingApp.NewService(nil)
	exportSvc := exportApp.NewService(searchStore)

	// Stub setup for Category templates seeding
	_ = typeStore

	kafkaConsumer := kafkaInfra.NewConsumer(permitStore)
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	mux := http.NewServeMux()
	httpIface.RegisterRoutes(mux, permitSvc, approvalSvc, riskSvc, searchSvc, reportingSvc, exportSvc)

	pipelineHandler := InitRouter(mux)
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "Permit-to-Work Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down Permit-to-Work Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	eventListener.Stop(shutdownCtx)
	_ = srv.Shutdown(shutdownCtx)

	prahariLogger.Info(ctx, "Permit-to-Work Service exited gracefully.")
}
