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
	complianceApp "prahari/services/compliance/internal/application/compliance"
	obligationApp "prahari/services/compliance/internal/application/obligation"
	reviewApp "prahari/services/compliance/internal/application/review"
	evidenceApp "prahari/services/compliance/internal/application/evidence"
	searchApp "prahari/services/compliance/internal/application/search"
	reportingApp "prahari/services/compliance/internal/application/reporting"
	exportApp "prahari/services/compliance/internal/application/export"

	// Infrastructure adapters
	pgInfra "prahari/services/compliance/internal/infrastructure/postgres"
	redisInfra "prahari/services/compliance/internal/infrastructure/redis"
	kafkaInfra "prahari/services/compliance/internal/infrastructure/kafka"
	workflowInfra "prahari/services/compliance/internal/infrastructure/workflow"
	identityInfra "prahari/services/compliance/internal/infrastructure/identity"
	eventsIface "prahari/services/compliance/internal/interfaces/events"

	// Interfaces
	httpIface "prahari/services/compliance/internal/interfaces/http"
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

	tp, err := InitTelemetry(ctx, "compliance-service", cfg.Environment)
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

	complianceStore := pgInfra.NewComplianceStore(db)
	_ = pgInfra.NewObligationStore(db)
	_ = pgInfra.NewRegulationStore(db)
	_ = pgInfra.NewStandardStore(db)
	_ = pgInfra.NewRequirementStore(db)
	_ = pgInfra.NewEvidenceStore(db)
	_ = pgInfra.NewControlStore(db)
	_ = pgInfra.NewCertificationStore(db)
	_ = pgInfra.NewFindingStore(db)
	_ = pgInfra.NewActionStore(db)
	reviewStore := pgInfra.NewReviewStore(db)
	_ = pgInfra.NewCommentStore(db)
	timelineStore := pgInfra.NewTimelineStore(db)
	auditStore := pgInfra.NewAuditStore(db)

	_ = redisInfra.NewCache(redisClient)
	kafkaPublisher := kafkaInfra.NewPublisher(nil)

	// DI Injection
	complianceSvc := complianceApp.NewService(complianceStore, kafkaPublisher, auditStore, timelineStore, workflowClient)
	obligationSvc := obligationApp.NewService(nil)
	reviewSvc := reviewApp.NewService(reviewStore)
	evidenceSvc := evidenceApp.NewService(nil)
	searchSvc := searchApp.NewService(complianceStore)
	reportingSvc := reportingApp.NewService(nil)
	exportSvc := exportApp.NewService(complianceStore)

	kafkaConsumer := kafkaInfra.NewConsumer(complianceStore)
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	mux := http.NewServeMux()
	httpIface.RegisterRoutes(mux, complianceSvc, obligationSvc, reviewSvc, evidenceSvc, searchSvc, reportingSvc, exportSvc)

	pipelineHandler := InitRouter(mux)
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "Compliance Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down Compliance Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	eventListener.Stop(shutdownCtx)
	_ = srv.Shutdown(shutdownCtx)

	prahariLogger.Info(ctx, "Compliance Service exited gracefully.")
}
