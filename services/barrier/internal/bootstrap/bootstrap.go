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
	analyticsApp "prahari/services/barrier/internal/application/analytics"
	barrierApp "prahari/services/barrier/internal/application/barrier"
	exportApp "prahari/services/barrier/internal/application/export"
	impairmentApp "prahari/services/barrier/internal/application/impairment"
	integrityApp "prahari/services/barrier/internal/application/integrity"
	prooftestApp "prahari/services/barrier/internal/application/prooftest"
	reportingApp "prahari/services/barrier/internal/application/reporting"
	searchApp "prahari/services/barrier/internal/application/search"

	// Infrastructure adapters
	kafkaInfra "prahari/services/barrier/internal/infrastructure/kafka"
	maintInfra "prahari/services/barrier/internal/infrastructure/maintenance"
	pgInfra "prahari/services/barrier/internal/infrastructure/postgres"
	redisInfra "prahari/services/barrier/internal/infrastructure/redis"

	// Interfaces
	eventsIface "prahari/services/barrier/internal/interfaces/events"
	httpIface "prahari/services/barrier/internal/interfaces/http"
)

func RunApp() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := LoadConfig(ctx)
	if err != nil {
		panic(err)
	}

	_ = InitLogger(cfg.Environment)

	tp, err := InitTelemetry(ctx, "barrier-service", cfg.Environment)
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

	maintClient := maintInfra.NewClient(cfg.MaintenanceGrpcAddr)

	// App Services DI
	barrierSvc := barrierApp.NewService(store, kafkaPublisher)
	integritySvc := integrityApp.NewService(store, kafkaPublisher)
	prooftestSvc := prooftestApp.NewService(store, kafkaPublisher, maintClient)
	impairmentSvc := impairmentApp.NewService(store, kafkaPublisher)
	reportingSvc := reportingApp.NewService(store)
	analyticsSvc := analyticsApp.NewService(store)
	searchSvc := searchApp.NewService(store)
	exportSvc := exportApp.NewService(store)

	kafkaConsumer := kafkaInfra.NewConsumer()
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	mux := http.NewServeMux()
	handler := httpIface.NewHandler(barrierSvc, integritySvc, prooftestSvc, impairmentSvc, reportingSvc, analyticsSvc, searchSvc, exportSvc)
	httpIface.RegisterRoutes(mux, handler)

	pipelineHandler := InitRouter(mux)
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "Barrier Management Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down Barrier Management Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	_ = srv.Shutdown(shutdownCtx)
	prahariLogger.Info(ctx, "Barrier Management Service exited gracefully.")
}
