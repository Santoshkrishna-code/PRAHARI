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
	analyticsApp "prahari/services/esg/internal/application/analytics"
	carbonApp "prahari/services/esg/internal/application/carbon"
	disclosureApp "prahari/services/esg/internal/application/disclosure"
	exportApp "prahari/services/esg/internal/application/export"
	reportingApp "prahari/services/esg/internal/application/reporting"
	searchApp "prahari/services/esg/internal/application/search"
	sustainabilityApp "prahari/services/esg/internal/application/sustainability"

	// Infrastructure adapters
	pgInfra "prahari/services/esg/internal/infrastructure/postgres"
	redisInfra "prahari/services/esg/internal/infrastructure/redis"
	kafkaInfra "prahari/services/esg/internal/infrastructure/kafka"
	workflowInfra "prahari/services/esg/internal/infrastructure/workflow"
	eventsIface "prahari/services/esg/internal/interfaces/events"

	// Interfaces
	httpIface "prahari/services/esg/internal/interfaces/http"
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

	tp, err := InitTelemetry(ctx, "esg-service", cfg.Environment)
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

	store := pgInfra.NewStore(db)

	_ = redisInfra.NewCache(redisClient)
	kafkaPublisher := kafkaInfra.NewPublisher()

	// DI Injection
	sustSvc := sustainabilityApp.NewService(store, store, store)
	monSvc := reportingApp.NewService(store, kafkaPublisher, store, store)
	carbonSvc := carbonApp.NewService(store, kafkaPublisher, store, store)
	disclSvc := disclosureApp.NewService(store, kafkaPublisher, store, store)
	searchSvc := searchApp.NewService(store)
	reportSvc := analyticsApp.NewService(store)
	exportSvc := exportApp.NewService(store)

	_ = workflowClient

	kafkaConsumer := kafkaInfra.NewConsumer()
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	mux := http.NewServeMux()
	handler := httpIface.NewHandler(sustSvc, monSvc, carbonSvc, disclSvc, searchSvc, reportSvc, exportSvc)
	httpIface.RegisterRoutes(mux, handler)

	pipelineHandler := InitRouter(mux)
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "ESG & Sustainability Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down ESG & Sustainability Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	_ = srv.Shutdown(shutdownCtx)

	prahariLogger.Info(ctx, "ESG & Sustainability Service exited gracefully.")
}
