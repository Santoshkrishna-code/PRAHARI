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
	contractorApp "prahari/services/contractor/internal/application/contractor"
	onboardingApp "prahari/services/contractor/internal/application/onboarding"
	complianceApp "prahari/services/contractor/internal/application/compliance"
	siteaccessApp "prahari/services/contractor/internal/application/siteaccess"
	searchApp "prahari/services/contractor/internal/application/search"
	reportingApp "prahari/services/contractor/internal/application/reporting"
	exportApp "prahari/services/contractor/internal/application/export"

	// Infrastructure adapters
	pgInfra "prahari/services/contractor/internal/infrastructure/postgres"
	redisInfra "prahari/services/contractor/internal/infrastructure/redis"
	kafkaInfra "prahari/services/contractor/internal/infrastructure/kafka"
	workflowInfra "prahari/services/contractor/internal/infrastructure/workflow"
	identityInfra "prahari/services/contractor/internal/infrastructure/identity"
	permitInfra "prahari/services/contractor/internal/infrastructure/permit"
	eventsIface "prahari/services/contractor/internal/interfaces/events"

	// Interfaces
	httpIface "prahari/services/contractor/internal/interfaces/http"
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

	tp, err := InitTelemetry(ctx, "contractor-service", cfg.Environment)
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

	contractorStore := pgInfra.NewContractorStore(db)
	workerStore := pgInfra.NewWorkerStore(db)
	certStore := pgInfra.NewCertStore(db)
	trainingStore := pgInfra.NewTrainingStore(db)
	_ = pgInfra.NewMedicalStore(db)
	accessStore := pgInfra.NewAccessStore(db)
	_ = pgInfra.NewInsuranceStore(db)
	_ = pgInfra.NewDocStore(db)
	_ = pgInfra.NewCommentStore(db)
	timelineStore := pgInfra.NewTimelineStore(db)
	auditStore := pgInfra.NewAuditStore(db)

	_ = redisInfra.NewCache(redisClient)
	kafkaPublisher := kafkaInfra.NewPublisher(nil)

	// DI Injection
	contractorSvc := contractorApp.NewService(contractorStore, kafkaPublisher, auditStore, timelineStore, identityClient, workflowClient)
	onboardingSvc := onboardingApp.NewService(workerStore)
	complianceSvc := complianceApp.NewService(certStore, trainingStore)
	siteaccessSvc := siteaccessApp.NewService(accessStore)
	searchSvc := searchApp.NewService(contractorStore)
	reportingSvc := reportingApp.NewService(nil)
	exportSvc := exportApp.NewService(contractorStore)

	kafkaConsumer := kafkaInfra.NewConsumer(contractorStore)
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	mux := http.NewServeMux()
	httpIface.RegisterRoutes(mux, contractorSvc, onboardingSvc, complianceSvc, siteaccessSvc, searchSvc, reportingSvc, exportSvc)

	pipelineHandler := InitRouter(mux)
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "Contractor Management Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down Contractor Management Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	eventListener.Stop(shutdownCtx)
	_ = srv.Shutdown(shutdownCtx)

	prahariLogger.Info(ctx, "Contractor Management Service exited gracefully.")
}
