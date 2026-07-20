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
	environmentApp "prahari/services/environmental/internal/application/environment"
	exportApp "prahari/services/environmental/internal/application/export"
	monitoringApp "prahari/services/environmental/internal/application/monitoring"
	permitApp "prahari/services/environmental/internal/application/permit"
	reportingApp "prahari/services/environmental/internal/application/reporting"
	searchApp "prahari/services/environmental/internal/application/search"
	wasteApp "prahari/services/environmental/internal/application/waste"

	// Infrastructure adapters
	pgInfra "prahari/services/environmental/internal/infrastructure/postgres"
	redisInfra "prahari/services/environmental/internal/infrastructure/redis"
	kafkaInfra "prahari/services/environmental/internal/infrastructure/kafka"
	workflowInfra "prahari/services/environmental/internal/infrastructure/workflow"
	eventsIface "prahari/services/environmental/internal/interfaces/events"

	// Interfaces
	httpIface "prahari/services/environmental/internal/interfaces/http"
)

// RunApp coordinates system DI injection and runs servers.
func RunApp() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := LoadConfig(ctx)
	if err != nil {
		panic(err)
	}

	_ = InitLogger(cfg.Environment)

	tp, err := InitTelemetry(ctx, "environmental-service", cfg.Environment)
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
	envSvc := environmentApp.NewService(store, store, store)
	monSvc := monitoringApp.NewService(store, kafkaPublisher, store, store)
	permitSvc := permitApp.NewService(store, store, store)
	wasteSvc := wasteApp.NewService(store, kafkaPublisher, store, store)
	searchSvc := searchApp.NewService(store)
	reportingSvc := reportingApp.NewService(store)
	exportSvc := exportApp.NewService(store)

	_ = workflowClient

	kafkaConsumer := kafkaInfra.NewConsumer()
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	mux := http.NewServeMux()
	handler := httpIface.NewHandler(envSvc, monSvc, permitSvc, wasteSvc, searchSvc, reportingSvc, exportSvc)
	httpIface.RegisterRoutes(mux, handler)

	pipelineHandler := InitRouter(mux)
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "Environmental Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down Environmental Service...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()

	_ = srv.Shutdown(shutdownCtx)

	prahariLogger.Info(ctx, "Environmental Service exited gracefully.")
}
