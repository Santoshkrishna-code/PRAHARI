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
	analyticsApp "prahari/services/emergency/internal/application/analytics"
	drillApp "prahari/services/emergency/internal/application/drill"
	evacuationApp "prahari/services/emergency/internal/application/evacuation"
	exportApp "prahari/services/emergency/internal/application/export"
	planningApp "prahari/services/emergency/internal/application/planning"
	recoveryApp "prahari/services/emergency/internal/application/recovery"
	reportingApp "prahari/services/emergency/internal/application/reporting"
	responseApp "prahari/services/emergency/internal/application/response"
	searchApp "prahari/services/emergency/internal/application/search"

	// Infrastructure adapters
	kafkaInfra "prahari/services/emergency/internal/infrastructure/kafka"
	notifInfra "prahari/services/emergency/internal/infrastructure/notification"
	pgInfra "prahari/services/emergency/internal/infrastructure/postgres"
	redisInfra "prahari/services/emergency/internal/infrastructure/redis"

	// Interfaces
	eventsIface "prahari/services/emergency/internal/interfaces/events"
	httpIface "prahari/services/emergency/internal/interfaces/http"
)

func RunApp() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := LoadConfig(ctx)
	if err != nil {
		panic(err)
	}

	_ = InitLogger(cfg.Environment)

	tp, err := InitTelemetry(ctx, "emergency-service", cfg.Environment)
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

	notifClient := notifInfra.NewClient(cfg.NotificationGrpcAddr)

	// App Services DI
	planningSvc := planningApp.NewService(store)
	responseSvc := responseApp.NewService(store, kafkaPublisher, notifClient)
	evacuationSvc := evacuationApp.NewService(store, kafkaPublisher)
	drillSvc := drillApp.NewService(store)
	recoverySvc := recoveryApp.NewService(store, kafkaPublisher)
	reportingSvc := reportingApp.NewService(store)
	analyticsSvc := analyticsApp.NewService(store)
	searchSvc := searchApp.NewService(store)
	exportSvc := exportApp.NewService(store)

	kafkaConsumer := kafkaInfra.NewConsumer()
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	mux := http.NewServeMux()
	handler := httpIface.NewHandler(planningSvc, responseSvc, evacuationSvc, drillSvc, recoverySvc, reportingSvc, analyticsSvc, searchSvc, exportSvc)
	httpIface.RegisterRoutes(mux, handler)

	pipelineHandler := InitRouter(mux)
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "Emergency Management Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down Emergency Management Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	_ = srv.Shutdown(shutdownCtx)
	prahariLogger.Info(ctx, "Emergency Management Service exited gracefully.")
}
