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
	analyticsApp "prahari/services/document/internal/application/analytics"
	approvalApp "prahari/services/document/internal/application/approval"
	creationApp "prahari/services/document/internal/application/creation"
	distributionApp "prahari/services/document/internal/application/distribution"
	exportApp "prahari/services/document/internal/application/export"
	lifecycleApp "prahari/services/document/internal/application/lifecycle"
	reportingApp "prahari/services/document/internal/application/reporting"
	searchApp "prahari/services/document/internal/application/search"
	versioningApp "prahari/services/document/internal/application/versioning"

	// Infrastructure adapters
	kafkaInfra "prahari/services/document/internal/infrastructure/kafka"
	pgInfra "prahari/services/document/internal/infrastructure/postgres"
	redisInfra "prahari/services/document/internal/infrastructure/redis"

	// Interfaces
	eventsIface "prahari/services/document/internal/interfaces/events"
	httpIface "prahari/services/document/internal/interfaces/http"
)

func RunApp() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := LoadConfig(ctx)
	if err != nil {
		panic(err)
	}

	_ = InitLogger(cfg.Environment)

	tp, err := InitTelemetry(ctx, "document-service", cfg.Environment)
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

	store := pgInfra.NewStore(db)
	_ = redisInfra.NewCache(redisClient)
	kafkaPublisher := kafkaInfra.NewPublisher()

	// App Services DI
	creationSvc := creationApp.NewService(store, kafkaPublisher)
	lifecycleSvc := lifecycleApp.NewService(store, kafkaPublisher)
	versioningSvc := versioningApp.NewService(store, kafkaPublisher)
	approvalSvc := approvalApp.NewService(store, kafkaPublisher)
	distributionSvc := distributionApp.NewService(store)
	reportingSvc := reportingApp.NewService(store)
	analyticsSvc := analyticsApp.NewService(store)
	searchSvc := searchApp.NewService(store)
	exportSvc := exportApp.NewService(store)

	kafkaConsumer := kafkaInfra.NewConsumer()
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	mux := http.NewServeMux()
	handler := httpIface.NewHandler(creationSvc, lifecycleSvc, versioningSvc, approvalSvc, distributionSvc, reportingSvc, analyticsSvc, searchSvc, exportSvc)
	httpIface.RegisterRoutes(mux, handler)

	pipelineHandler := InitRouter(mux)
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "Document Management Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down Document Management Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	_ = srv.Shutdown(shutdownCtx)
	prahariLogger.Info(ctx, "Document Management Service exited gracefully.")
}
