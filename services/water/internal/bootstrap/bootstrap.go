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
	analyticsApp "prahari/services/water/internal/application/analytics"
	consumptionApp "prahari/services/water/internal/application/consumption"
	exportApp "prahari/services/water/internal/application/export"
	forecastingApp "prahari/services/water/internal/application/forecasting"
	optimizationApp "prahari/services/water/internal/application/optimization"
	recyclingApp "prahari/services/water/internal/application/recycling"
	reportingApp "prahari/services/water/internal/application/reporting"
	searchApp "prahari/services/water/internal/application/search"

	// Infrastructure adapters
	energyInfra "prahari/services/water/internal/infrastructure/energy"
	envInfra "prahari/services/water/internal/infrastructure/environmental"
	esgInfra "prahari/services/water/internal/infrastructure/esg"
	kafkaInfra "prahari/services/water/internal/infrastructure/kafka"
	maintInfra "prahari/services/water/internal/infrastructure/maintenance"
	pgInfra "prahari/services/water/internal/infrastructure/postgres"
	redisInfra "prahari/services/water/internal/infrastructure/redis"

	// Interfaces
	eventsIface "prahari/services/water/internal/interfaces/events"
	httpIface "prahari/services/water/internal/interfaces/http"
)

func RunApp() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := LoadConfig(ctx)
	if err != nil {
		panic(err)
	}

	_ = InitLogger(cfg.Environment)

	tp, err := InitTelemetry(ctx, "water-service", cfg.Environment)
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

	envClient := envInfra.NewClient(cfg.EnvironmentalGrpcAddr)
	esgClient := esgInfra.NewClient(cfg.ESGGrpcAddr)
	maintClient := maintInfra.NewClient(cfg.MaintenanceGrpcAddr)
	_ = energyInfra.NewClient(cfg.EnergyGrpcAddr)

	// App Services DI
	reportingSvc := reportingApp.NewService(store, kafkaPublisher)
	consumptionSvc := consumptionApp.NewService(store, kafkaPublisher, envClient, esgClient)
	recyclingSvc := recyclingApp.NewService(store, kafkaPublisher)
	forecastingSvc := forecastingApp.NewService(store, kafkaPublisher)
	optimizationSvc := optimizationApp.NewService(store)
	_ = distributionNetworkAppInit(store, maintClient)
	searchSvc := searchApp.NewService(store)
	analyticsSvc := analyticsApp.NewService(store)
	exportSvc := exportApp.NewService(store)

	kafkaConsumer := kafkaInfra.NewConsumer()
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	mux := http.NewServeMux()
	handler := httpIface.NewHandler(reportingSvc, consumptionSvc, recyclingSvc, forecastingSvc, optimizationSvc, searchSvc, analyticsSvc, exportSvc)
	httpIface.RegisterRoutes(mux, handler)

	pipelineHandler := InitRouter(mux)
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "Water Management Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down Water Management Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	_ = srv.Shutdown(shutdownCtx)
	prahariLogger.Info(ctx, "Water Management Service exited gracefully.")
}

func distributionNetworkAppInit(store *pgInfra.Store, maintClient *maintInfra.Client) any {
	return nil
}
