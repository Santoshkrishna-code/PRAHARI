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
	hazardApp "prahari/services/hazard/internal/application/hazard"
	assessmentApp "prahari/services/hazard/internal/application/assessment"
	mitigationApp "prahari/services/hazard/internal/application/mitigation"
	verifyApp "prahari/services/hazard/internal/application/verification"
	searchApp "prahari/services/hazard/internal/application/search"
	reportingApp "prahari/services/hazard/internal/application/reporting"
	exportApp "prahari/services/hazard/internal/application/export"

	// Infrastructure adapters
	pgInfra "prahari/services/hazard/internal/infrastructure/postgres"
	redisInfra "prahari/services/hazard/internal/infrastructure/redis"
	kafkaInfra "prahari/services/hazard/internal/infrastructure/kafka"
	workflowInfra "prahari/services/hazard/internal/infrastructure/workflow"
	identityInfra "prahari/services/hazard/internal/infrastructure/identity"
	incidentInfra "prahari/services/hazard/internal/infrastructure/incident"
	eventsIface "prahari/services/hazard/internal/interfaces/events"

	// Interfaces
	httpIface "prahari/services/hazard/internal/interfaces/http"
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

	tp, err := InitTelemetry(ctx, "hazard-service", cfg.Environment)
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
	incidentClient := incidentInfra.NewClient(cfg.IncidentGrpcAddr)

	hazardStore := pgInfra.NewHazardStore(db)
	_ = pgInfra.NewTypeStore(db)
	_ = pgInfra.NewCategoryStore(db)
	_ = pgInfra.NewLocationStore(db)
	controlStore := pgInfra.NewControlStore(db)
	mitigationStore := pgInfra.NewMitigationStore(db)
	assessmentStore := pgInfra.NewAssessmentStore(db)
	verifyStore := pgInfra.NewVerifyStore(db)
	_ = pgInfra.NewCommentStore(db)
	timelineStore := pgInfra.NewTimelineStore(db)
	auditStore := pgInfra.NewAuditStore(db)

	_ = redisInfra.NewCache(redisClient)
	kafkaPublisher := kafkaInfra.NewPublisher(nil)

	// DI Injection
	hazardSvc := hazardApp.NewService(hazardStore, kafkaPublisher, auditStore, timelineStore, incidentClient, workflowClient)
	assessmentSvc := assessmentApp.NewService(assessmentStore)
	mitigationSvc := mitigationApp.NewService(mitigationStore)
	verifySvc := verifyApp.NewService(verifyStore)
	searchSvc := searchApp.NewService(hazardStore)
	reportingSvc := reportingApp.NewService(nil)
	exportSvc := exportApp.NewService(hazardStore)

	kafkaConsumer := kafkaInfra.NewConsumer(hazardStore)
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	mux := http.NewServeMux()
	httpIface.RegisterRoutes(mux, hazardSvc, assessmentSvc, mitigationSvc, verifySvc, searchSvc, reportingSvc, exportSvc)

	pipelineHandler := InitRouter(mux)
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "Hazard Management Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down Hazard Management Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	eventListener.Stop(shutdownCtx)
	_ = srv.Shutdown(shutdownCtx)

	prahariLogger.Info(ctx, "Hazard Management Service exited gracefully.")
}
