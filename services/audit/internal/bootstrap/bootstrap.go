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
	auditApp "prahari/services/audit/internal/application/audit"
	planningApp "prahari/services/audit/internal/application/planning"
	executionApp "prahari/services/audit/internal/application/execution"
	reviewApp "prahari/services/audit/internal/application/review"
	findingsApp "prahari/services/audit/internal/application/findings"
	searchApp "prahari/services/audit/internal/application/search"
	reportingApp "prahari/services/audit/internal/application/reporting"
	exportApp "prahari/services/audit/internal/application/export"

	// Infrastructure adapters
	pgInfra "prahari/services/audit/internal/infrastructure/postgres"
	redisInfra "prahari/services/audit/internal/infrastructure/redis"
	kafkaInfra "prahari/services/audit/internal/infrastructure/kafka"
	workflowInfra "prahari/services/audit/internal/infrastructure/workflow"
	identityInfra "prahari/services/audit/internal/infrastructure/identity"
	eventsIface "prahari/services/audit/internal/interfaces/events"

	// Interfaces
	httpIface "prahari/services/audit/internal/interfaces/http"
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

	tp, err := InitTelemetry(ctx, "audit-service", cfg.Environment)
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

	auditStore := pgInfra.NewAuditStore(db)
	_ = pgInfra.NewProgramStore(db)
	_ = pgInfra.NewPlanStore(db)
	_ = pgInfra.NewTypeStore(db)
	_ = pgInfra.NewChecklistStore(db)
	_ = pgInfra.NewItemStore(db)
	_ = pgInfra.NewAuditorStore(db)
	_ = pgInfra.NewAuditeeStore(db)
	_ = pgInfra.NewFindingStore(db)
	_ = pgInfra.NewNCRStore(db)
	_ = pgInfra.NewActionStore(db)
	_ = pgInfra.NewEvidenceStore(db)
	reviewStore := pgInfra.NewReviewStore(db)
	_ = pgInfra.NewCommentStore(db)
	timelineStore := pgInfra.NewTimelineStore(db)
	trailStore := pgInfra.NewTrailStore(db)

	_ = redisInfra.NewCache(redisClient)
	kafkaPublisher := kafkaInfra.NewPublisher(nil)

	// DI Injection
	auditSvc := auditApp.NewService(auditStore, kafkaPublisher, trailStore, timelineStore, workflowClient)
	planningSvc := planningApp.NewService(nil)
	executionSvc := executionApp.NewService(nil)
	reviewSvc := reviewApp.NewService(reviewStore)
	findingsSvc := findingsApp.NewService(nil)
	searchSvc := searchApp.NewService(auditStore)
	reportingSvc := reportingApp.NewService(nil)
	exportSvc := exportApp.NewService(auditStore)

	kafkaConsumer := kafkaInfra.NewConsumer(auditStore)
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	mux := http.NewServeMux()
	httpIface.RegisterRoutes(mux, auditSvc, planningSvc, executionSvc, reviewSvc, findingsSvc, searchSvc, reportingSvc, exportSvc)

	pipelineHandler := InitRouter(mux)
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "Audit Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down Audit Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	eventListener.Stop(shutdownCtx)
	_ = srv.Shutdown(shutdownCtx)

	prahariLogger.Info(ctx, "Audit Service exited gracefully.")
}
