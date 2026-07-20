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
	analyticsApp "prahari/services/loto/internal/application/analytics"
	executionApp "prahari/services/loto/internal/application/execution"
	exportApp "prahari/services/loto/internal/application/export"
	planningApp "prahari/services/loto/internal/application/planning"
	reportingApp "prahari/services/loto/internal/application/reporting"
	restorationApp "prahari/services/loto/internal/application/restoration"
	searchApp "prahari/services/loto/internal/application/search"
	verificationApp "prahari/services/loto/internal/application/verification"

	// Infrastructure adapters
	kafkaInfra "prahari/services/loto/internal/infrastructure/kafka"
	pgInfra "prahari/services/loto/internal/infrastructure/postgres"
	redisInfra "prahari/services/loto/internal/infrastructure/redis"

	// Interfaces
	eventsIface "prahari/services/loto/internal/interfaces/events"
	httpIface "prahari/services/loto/internal/interfaces/http"
)

func RunApp() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := LoadConfig(ctx)
	if err != nil {
		panic(err)
	}

	_ = InitLogger(cfg.Environment)

	tp, err := InitTelemetry(ctx, "loto-service", cfg.Environment)
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
	planningSvc := planningApp.NewService(store, kafkaPublisher)
	executionSvc := executionApp.NewService(store, kafkaPublisher)
	verificationSvc := verificationApp.NewService(store, kafkaPublisher)
	restorationSvc := restorationApp.NewService(store, kafkaPublisher)
	reportingSvc := reportingApp.NewService(store)
	analyticsSvc := analyticsApp.NewService(store)
	searchSvc := searchApp.NewService(store)
	exportSvc := exportApp.NewService(store)

	kafkaConsumer := kafkaInfra.NewConsumer()
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	mux := http.NewServeMux()
	handler := httpIface.NewHandler(planningSvc, executionSvc, verificationSvc, restorationSvc, reportingSvc, analyticsSvc, searchSvc, exportSvc)
	httpIface.RegisterRoutes(mux, handler)

	pipelineHandler := InitRouter(mux)
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "LOTO & Isolation Management Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down LOTO Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	_ = srv.Shutdown(shutdownCtx)
	prahariLogger.Info(ctx, "LOTO Service exited gracefully.")
}
