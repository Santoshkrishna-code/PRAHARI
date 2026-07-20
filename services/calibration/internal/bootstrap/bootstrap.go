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
	analyticsApp "prahari/services/calibration/internal/application/analytics"
	certificateApp "prahari/services/calibration/internal/application/certificate"
	executionApp "prahari/services/calibration/internal/application/execution"
	exportApp "prahari/services/calibration/internal/application/export"
	reportingApp "prahari/services/calibration/internal/application/reporting"
	schedulingApp "prahari/services/calibration/internal/application/scheduling"
	searchApp "prahari/services/calibration/internal/application/search"

	// Infrastructure adapters
	kafkaInfra "prahari/services/calibration/internal/infrastructure/kafka"
	pgInfra "prahari/services/calibration/internal/infrastructure/postgres"
	redisInfra "prahari/services/calibration/internal/infrastructure/redis"

	// Interfaces
	eventsIface "prahari/services/calibration/internal/interfaces/events"
	httpIface "prahari/services/calibration/internal/interfaces/http"
)

func RunApp() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := LoadConfig(ctx)
	if err != nil {
		panic(err)
	}

	_ = InitLogger(cfg.Environment)

	tp, err := InitTelemetry(ctx, "calibration-service", cfg.Environment)
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
	schedulingSvc := schedulingApp.NewService(store, kafkaPublisher)
	executionSvc := executionApp.NewService(store, kafkaPublisher)
	certificateSvc := certificateApp.NewService(store, kafkaPublisher)
	reportingSvc := reportingApp.NewService(store)
	analyticsSvc := analyticsApp.NewService(store)
	searchSvc := searchApp.NewService(store)
	exportSvc := exportApp.NewService(store)

	kafkaConsumer := kafkaInfra.NewConsumer()
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	mux := http.NewServeMux()
	handler := httpIface.NewHandler(schedulingSvc, executionSvc, certificateSvc, reportingSvc, analyticsSvc, searchSvc, exportSvc)
	httpIface.RegisterRoutes(mux, handler)

	pipelineHandler := InitRouter(mux)
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "Calibration Management Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down Calibration Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	_ = srv.Shutdown(shutdownCtx)
	prahariLogger.Info(ctx, "Calibration Service exited gracefully.")
}
