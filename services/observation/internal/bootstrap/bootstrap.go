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
	observationApp "prahari/services/observation/internal/application/observation"
	coachingApp "prahari/services/observation/internal/application/coaching"
	followupApp "prahari/services/observation/internal/application/followup"
	effectivenessApp "prahari/services/observation/internal/application/effectiveness"
	searchApp "prahari/services/observation/internal/application/search"
	reportingApp "prahari/services/observation/internal/application/reporting"
	exportApp "prahari/services/observation/internal/application/export"

	// Infrastructure adapters
	pgInfra "prahari/services/observation/internal/infrastructure/postgres"
	redisInfra "prahari/services/observation/internal/infrastructure/redis"
	kafkaInfra "prahari/services/observation/internal/infrastructure/kafka"
	workflowInfra "prahari/services/observation/internal/infrastructure/workflow"
	identityInfra "prahari/services/observation/internal/infrastructure/identity"
	hazardInfra "prahari/services/observation/internal/infrastructure/hazard"
	incidentInfra "prahari/services/observation/internal/infrastructure/incident"
	eventsIface "prahari/services/observation/internal/interfaces/events"

	// Interfaces
	httpIface "prahari/services/observation/internal/interfaces/http"
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

	tp, err := InitTelemetry(ctx, "observation-service", cfg.Environment)
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
	hazardClient := hazardInfra.NewClient(cfg.HazardGrpcAddr)
	incidentClient := incidentInfra.NewClient(cfg.IncidentGrpcAddr)

	observationStore := pgInfra.NewObservationStore(db)
	_ = pgInfra.NewTypeStore(db)
	_ = pgInfra.NewCategoryStore(db)
	coachingStore := pgInfra.NewCoachingStore(db)
	_ = pgInfra.NewFollowUpStore(db)
	_ = pgInfra.NewEffectivenessStore(db)
	_ = pgInfra.NewActionStore(db)
	_ = pgInfra.NewCommentStore(db)
	timelineStore := pgInfra.NewTimelineStore(db)
	auditStore := pgInfra.NewAuditStore(db)

	_ = redisInfra.NewCache(redisClient)
	kafkaPublisher := kafkaInfra.NewPublisher(nil)

	// DI Injection
	observationSvc := observationApp.NewService(observationStore, kafkaPublisher, auditStore, timelineStore, hazardClient, incidentClient)
	coachingSvc := coachingApp.NewService(coachingStore)
	followupSvc := followupApp.NewService(nil)
	effectivenessSvc := effectivenessApp.NewService(nil)
	searchSvc := searchApp.NewService(observationStore)
	reportingSvc := reportingApp.NewService(nil)
	exportSvc := exportApp.NewService(observationStore)

	kafkaConsumer := kafkaInfra.NewConsumer(observationStore)
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	mux := http.NewServeMux()
	httpIface.RegisterRoutes(mux, observationSvc, coachingSvc, followupSvc, effectivenessSvc, searchSvc, reportingSvc, exportSvc)

	pipelineHandler := InitRouter(mux)
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "Observation Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down Observation Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	eventListener.Stop(shutdownCtx)
	_ = srv.Shutdown(shutdownCtx)

	prahariLogger.Info(ctx, "Observation Service exited gracefully.")
}
