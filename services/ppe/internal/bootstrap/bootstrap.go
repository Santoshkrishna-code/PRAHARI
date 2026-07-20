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
	analyticsApp "prahari/services/ppe/internal/application/analytics"
	catalogApp "prahari/services/ppe/internal/application/catalog"
	exportApp "prahari/services/ppe/internal/application/export"
	inspectionApp "prahari/services/ppe/internal/application/inspection"
	issuanceApp "prahari/services/ppe/internal/application/issuance"
	maintenanceApp "prahari/services/ppe/internal/application/maintenance"
	reportingApp "prahari/services/ppe/internal/application/reporting"
	searchApp "prahari/services/ppe/internal/application/search"

	// Infrastructure adapters
	kafkaInfra "prahari/services/ppe/internal/infrastructure/kafka"
	pgInfra "prahari/services/ppe/internal/infrastructure/postgres"
	redisInfra "prahari/services/ppe/internal/infrastructure/redis"

	// Interfaces
	eventsIface "prahari/services/ppe/internal/interfaces/events"
	httpIface "prahari/services/ppe/internal/interfaces/http"
)

func RunApp() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := LoadConfig(ctx)
	if err != nil {
		panic(err)
	}

	_ = InitLogger(cfg.Environment)

	tp, err := InitTelemetry(ctx, "ppe-service", cfg.Environment)
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
	catalogSvc := catalogApp.NewService(store, kafkaPublisher)
	issuanceSvc := issuanceApp.NewService(store, kafkaPublisher)
	inspectionSvc := inspectionApp.NewService(store, kafkaPublisher)
	maintenanceSvc := maintenanceApp.NewService(store, kafkaPublisher)
	reportingSvc := reportingApp.NewService(store)
	analyticsSvc := analyticsApp.NewService(store)
	searchSvc := searchApp.NewService(store)
	exportSvc := exportApp.NewService(store)

	kafkaConsumer := kafkaInfra.NewConsumer()
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	mux := http.NewServeMux()
	handler := httpIface.NewHandler(catalogSvc, issuanceSvc, inspectionSvc, maintenanceSvc, reportingSvc, analyticsSvc, searchSvc, exportSvc)
	httpIface.RegisterRoutes(mux, handler)

	pipelineHandler := InitRouter(mux)
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "PPE Management Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down PPE Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	_ = srv.Shutdown(shutdownCtx)
	prahariLogger.Info(ctx, "PPE Service exited gracefully.")
}
