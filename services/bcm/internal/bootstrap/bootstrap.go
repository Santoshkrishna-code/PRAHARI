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
	analyticsApp "prahari/services/bcm/internal/application/analytics"
	biaApp "prahari/services/bcm/internal/application/bia"
	exercisesApp "prahari/services/bcm/internal/application/exercises"
	exportApp "prahari/services/bcm/internal/application/export"
	planningApp "prahari/services/bcm/internal/application/planning"
	recoveryApp "prahari/services/bcm/internal/application/recovery"
	reportingApp "prahari/services/bcm/internal/application/reporting"
	resilienceApp "prahari/services/bcm/internal/application/resilience"
	searchApp "prahari/services/bcm/internal/application/search"

	// Infrastructure adapters
	kafkaInfra "prahari/services/bcm/internal/infrastructure/kafka"
	pgInfra "prahari/services/bcm/internal/infrastructure/postgres"
	redisInfra "prahari/services/bcm/internal/infrastructure/redis"

	// Interfaces
	eventsIface "prahari/services/bcm/internal/interfaces/events"
	httpIface "prahari/services/bcm/internal/interfaces/http"
)

func RunApp() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := LoadConfig(ctx)
	if err != nil {
		panic(err)
	}

	_ = InitLogger(cfg.Environment)

	tp, err := InitTelemetry(ctx, "bcm-service", cfg.Environment)
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
	biaSvc := biaApp.NewService(store, kafkaPublisher)
	planningSvc := planningApp.NewService(store, kafkaPublisher)
	recoverySvc := recoveryApp.NewService(store, kafkaPublisher)
	exercisesSvc := exercisesApp.NewService(store)
	resilienceSvc := resilienceApp.NewService(store)
	reportingSvc := reportingApp.NewService(store)
	analyticsSvc := analyticsApp.NewService(store)
	searchSvc := searchApp.NewService(store)
	exportSvc := exportApp.NewService(store)

	kafkaConsumer := kafkaInfra.NewConsumer()
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	mux := http.NewServeMux()
	handler := httpIface.NewHandler(biaSvc, planningSvc, recoverySvc, exercisesSvc, resilienceSvc, reportingSvc, analyticsSvc, searchSvc, exportSvc)
	httpIface.RegisterRoutes(mux, handler)

	pipelineHandler := InitRouter(mux)
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "Business Continuity Management (BCM) Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down BCM Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	_ = srv.Shutdown(shutdownCtx)
	prahariLogger.Info(ctx, "BCM Service exited gracefully.")
}
