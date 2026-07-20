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
	clearanceApp "prahari/services/occupational-health/internal/application/clearance"
	exportApp "prahari/services/occupational-health/internal/application/export"
	medicalApp "prahari/services/occupational-health/internal/application/medical"
	reportingApp "prahari/services/occupational-health/internal/application/reporting"
	searchApp "prahari/services/occupational-health/internal/application/search"
	exposureApp "prahari/services/occupational-health/internal/application/exposure"
	surveillanceApp "prahari/services/occupational-health/internal/application/surveillance"
	vaccinationApp "prahari/services/occupational-health/internal/application/vaccination"

	// Infrastructure adapters
	pgInfra "prahari/services/occupational-health/internal/infrastructure/postgres"
	redisInfra "prahari/services/occupational-health/internal/infrastructure/redis"
	kafkaInfra "prahari/services/occupational-health/internal/infrastructure/kafka"
	workflowInfra "prahari/services/occupational-health/internal/infrastructure/workflow"
	eventsIface "prahari/services/occupational-health/internal/interfaces/events"

	// Interfaces
	httpIface "prahari/services/occupational-health/internal/interfaces/http"
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

	tp, err := InitTelemetry(ctx, "occupational-health-service", cfg.Environment)
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
	clearanceSvc := clearanceApp.NewService(store, kafkaPublisher, store, store, workflowClient)
	medicalSvc := medicalApp.NewService(store)
	searchSvc := searchApp.NewService(store)
	reportingSvc := reportingApp.NewService(store)
	exportSvc := exportApp.NewService(store)
	_ = exposureApp.NewService(store)
	_ = surveillanceApp.NewService(store)
	_ = vaccinationApp.NewService(store)


	kafkaConsumer := kafkaInfra.NewConsumer(clearanceSvc)
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	mux := http.NewServeMux()
	handler := httpIface.NewHandler(clearanceSvc, medicalSvc, searchSvc, reportingSvc, exportSvc)
	httpIface.RegisterRoutes(mux, handler)

	pipelineHandler := InitRouter(mux)
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "Occupational Health Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down Occupational Health Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	_ = srv.Shutdown(shutdownCtx)

	prahariLogger.Info(ctx, "Occupational Health Service exited gracefully.")
}
