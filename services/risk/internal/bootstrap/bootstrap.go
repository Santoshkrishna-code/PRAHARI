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
	assessmentApp "prahari/services/risk/internal/application/assessment"
	reviewApp "prahari/services/risk/internal/application/review"
	approvalApp "prahari/services/risk/internal/application/approval"
	residualApp "prahari/services/risk/internal/application/residual"
	searchApp "prahari/services/risk/internal/application/search"
	reportingApp "prahari/services/risk/internal/application/reporting"
	exportApp "prahari/services/risk/internal/application/export"

	// Infrastructure adapters
	pgInfra "prahari/services/risk/internal/infrastructure/postgres"
	redisInfra "prahari/services/risk/internal/infrastructure/redis"
	kafkaInfra "prahari/services/risk/internal/infrastructure/kafka"
	workflowInfra "prahari/services/risk/internal/infrastructure/workflow"
	identityInfra "prahari/services/risk/internal/infrastructure/identity"
	eventsIface "prahari/services/risk/internal/interfaces/events"

	// Interfaces
	httpIface "prahari/services/risk/internal/interfaces/http"
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

	tp, err := InitTelemetry(ctx, "risk-service", cfg.Environment)
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

	assessmentStore := pgInfra.NewAssessmentStore(db)
	_ = pgInfra.NewRegisterStore(db)
	_ = pgInfra.NewControlStore(db)
	_ = pgInfra.NewBarrierStore(db)
	_ = pgInfra.NewMatrixStore(db)
	reviewStore := pgInfra.NewReviewStore(db)
	approvalStore := pgInfra.NewApprovalStore(db)
	residualStore := pgInfra.NewResidualStore(db)
	_ = pgInfra.NewCommentStore(db)
	timelineStore := pgInfra.NewTimelineStore(db)
	auditStore := pgInfra.NewAuditStore(db)

	_ = redisInfra.NewCache(redisClient)
	kafkaPublisher := kafkaInfra.NewPublisher(nil)

	// DI Injection
	assessmentSvc := assessmentApp.NewService(assessmentStore, kafkaPublisher, auditStore, timelineStore, workflowClient)
	reviewSvc := reviewApp.NewService(reviewStore)
	approvalSvc := approvalApp.NewService(approvalStore)
	residualSvc := residualApp.NewService(residualStore)
	searchSvc := searchApp.NewService(assessmentStore)
	reportingSvc := reportingApp.NewService(nil)
	exportSvc := exportApp.NewService(assessmentStore)

	kafkaConsumer := kafkaInfra.NewConsumer(assessmentStore)
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	mux := http.NewServeMux()
	httpIface.RegisterRoutes(mux, assessmentSvc, reviewSvc, approvalSvc, residualSvc, searchSvc, reportingSvc, exportSvc)

	pipelineHandler := InitRouter(mux)
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "Risk Assessment Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down Risk Assessment Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	eventListener.Stop(shutdownCtx)
	_ = srv.Shutdown(shutdownCtx)

	prahariLogger.Info(ctx, "Risk Assessment Service exited gracefully.")
}
