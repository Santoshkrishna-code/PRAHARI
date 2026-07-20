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
	nearmissApp "prahari/services/nearmiss/internal/application/nearmiss"
	investigationApp "prahari/services/nearmiss/internal/application/investigation"
	correctiveApp "prahari/services/nearmiss/internal/application/corrective"
	verifyApp "prahari/services/nearmiss/internal/application/verification"
	searchApp "prahari/services/nearmiss/internal/application/search"
	reportingApp "prahari/services/nearmiss/internal/application/reporting"
	exportApp "prahari/services/nearmiss/internal/application/export"

	// Infrastructure adapters
	pgInfra "prahari/services/nearmiss/internal/infrastructure/postgres"
	redisInfra "prahari/services/nearmiss/internal/infrastructure/redis"
	kafkaInfra "prahari/services/nearmiss/internal/infrastructure/kafka"
	workflowInfra "prahari/services/nearmiss/internal/infrastructure/workflow"
	identityInfra "prahari/services/nearmiss/internal/infrastructure/identity"
	incidentInfra "prahari/services/nearmiss/internal/infrastructure/incident"
	hazardInfra "prahari/services/nearmiss/internal/infrastructure/hazard"
	eventsIface "prahari/services/nearmiss/internal/interfaces/events"

	// Interfaces
	httpIface "prahari/services/nearmiss/internal/interfaces/http"
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

	tp, err := InitTelemetry(ctx, "nearmiss-service", cfg.Environment)
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
	hazardClient := hazardInfra.NewClient(cfg.HazardGrpcAddr)

	nearmissStore := pgInfra.NewNearMissStore(db)
	_ = pgInfra.NewClassStore(db)
	_ = pgInfra.NewSeverityStore(db)
	_ = pgInfra.NewCauseStore(db)
	investigationStore := pgInfra.NewInvestigationStore(db)
	correctiveStore := pgInfra.NewCorrectiveStore(db)
	_ = pgInfra.NewPreventiveStore(db)
	verifyStore := pgInfra.NewVerifyStore(db)
	_ = pgInfra.NewCommentStore(db)
	timelineStore := pgInfra.NewTimelineStore(db)
	auditStore := pgInfra.NewAuditStore(db)

	_ = redisInfra.NewCache(redisClient)
	kafkaPublisher := kafkaInfra.NewPublisher(nil)

	// DI Injection
	nearmissSvc := nearmissApp.NewService(nearmissStore, kafkaPublisher, auditStore, timelineStore, incidentClient, hazardClient)
	investigationSvc := investigationApp.NewService(investigationStore)
	correctiveSvc := correctiveApp.NewService(correctiveStore)
	verifySvc := verifyApp.NewService(verifyStore)
	searchSvc := searchApp.NewService(nearmissStore)
	reportingSvc := reportingApp.NewService(nil)
	exportSvc := exportApp.NewService(nearmissStore)

	kafkaConsumer := kafkaInfra.NewConsumer(nearmissStore)
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	mux := http.NewServeMux()
	httpIface.RegisterRoutes(mux, nearmissSvc, investigationSvc, correctiveSvc, verifySvc, searchSvc, reportingSvc, exportSvc)

	pipelineHandler := InitRouter(mux)
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "Near Miss Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down Near Miss Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	eventListener.Stop(shutdownCtx)
	_ = srv.Shutdown(shutdownCtx)

	prahariLogger.Info(ctx, "Near Miss Service exited gracefully.")
}
