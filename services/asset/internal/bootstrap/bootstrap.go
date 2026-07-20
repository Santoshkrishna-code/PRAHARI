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
	assetApp "prahari/services/asset/internal/application/asset"
	lifecycleApp "prahari/services/asset/internal/application/lifecycle"
	searchApp "prahari/services/asset/internal/application/search"
	reportingApp "prahari/services/asset/internal/application/reporting"
	exportApp "prahari/services/asset/internal/application/export"

	// Infrastructure adapters
	pgInfra "prahari/services/asset/internal/infrastructure/postgres"
	redisInfra "prahari/services/asset/internal/infrastructure/redis"
	kafkaInfra "prahari/services/asset/internal/infrastructure/kafka"
	workflowInfra "prahari/services/asset/internal/infrastructure/workflow"
	identityInfra "prahari/services/asset/internal/infrastructure/identity"
	inspectionInfra "prahari/services/asset/internal/infrastructure/inspection"
	eventsIface "prahari/services/asset/internal/interfaces/events"

	// Interfaces
	httpIface "prahari/services/asset/internal/interfaces/http"
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

	tp, err := InitTelemetry(ctx, "asset-service", cfg.Environment)
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
	inspectionClient := inspectionInfra.NewClient(cfg.InspectionGrpcAddr)

	assetStore := pgInfra.NewAssetStore(db)
	_ = pgInfra.NewHierarchyStore(db)
	_ = pgInfra.NewRelationshipStore(db)
	_ = pgInfra.NewTypeStore(db)
	_ = pgInfra.NewCategoryStore(db)
	_ = pgInfra.NewLocationStore(db)
	_ = pgInfra.NewManufacturerStore(db)
	_ = pgInfra.NewSpecStore(db)
	_ = pgInfra.NewWarrantyStore(db)
	_ = pgInfra.NewDocStore(db)
	_ = pgInfra.NewCommentStore(db)
	timelineStore := pgInfra.NewTimelineStore(db)
	auditStore := pgInfra.NewAuditStore(db)

	_ = redisInfra.NewCache(redisClient)
	kafkaPublisher := kafkaInfra.NewPublisher(nil)

	// DI Injection
	assetSvc := assetApp.NewService(assetStore, kafkaPublisher, auditStore, timelineStore, inspectionClient, workflowClient)
	lifecycleSvc := lifecycleApp.NewService(timelineStore)
	searchSvc := searchApp.NewService(assetStore)
	reportingSvc := reportingApp.NewService(nil)
	exportSvc := exportApp.NewService(assetStore)

	kafkaConsumer := kafkaInfra.NewConsumer(assetStore)
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	mux := http.NewServeMux()
	httpIface.RegisterRoutes(mux, assetSvc, lifecycleSvc, searchSvc, reportingSvc, exportSvc)

	pipelineHandler := InitRouter(mux)
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "Asset Management Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down Asset Management Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	eventListener.Stop(shutdownCtx)
	_ = srv.Shutdown(shutdownCtx)

	prahariLogger.Info(ctx, "Asset Management Service exited gracefully.")
}
