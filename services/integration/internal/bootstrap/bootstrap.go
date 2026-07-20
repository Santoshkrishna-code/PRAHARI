package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"
	"go.opentelemetry.io/otel/sdk/trace"

	prahariLogger "prahari/shared/logger"
	prahariMid "prahari/shared/middleware"
	prahariCORS "prahari/shared/middleware/cors"
	prahariLogMid "prahari/shared/middleware/logging"
	prahariRecovery "prahari/shared/middleware/recovery"
	prahariRequest "prahari/shared/middleware/requestid"
	prahariTimeout "prahari/shared/middleware/timeout"
	prahariRedis "prahari/shared/redis"
	prahariTelemetry "prahari/shared/telemetry"
	prahariExporters "prahari/shared/telemetry/exporters"

	// Application services
	analyticsApp "prahari/services/integration/internal/application/analytics"
	connectorsApp "prahari/services/integration/internal/application/connectors"
	exportApp "prahari/services/integration/internal/application/export"
	reportingApp "prahari/services/integration/internal/application/reporting"
	schedulingApp "prahari/services/integration/internal/application/scheduling"
	searchApp "prahari/services/integration/internal/application/search"
	synchronizationApp "prahari/services/integration/internal/application/synchronization"
	transformationApp "prahari/services/integration/internal/application/transformation"

	// Infrastructure adapters
	kafkaInfra "prahari/services/integration/internal/infrastructure/kafka"
	mqttInfra "prahari/services/integration/internal/infrastructure/mqtt"
	pgInfra "prahari/services/integration/internal/infrastructure/postgres"
	redisInfra "prahari/services/integration/internal/infrastructure/redis"

	// Interfaces
	eventsIface "prahari/services/integration/internal/interfaces/events"
	grpcIface "prahari/services/integration/internal/interfaces/grpc"
	httpIface "prahari/services/integration/internal/interfaces/http"
)

func initTelemetry(ctx context.Context, env string) (*trace.TracerProvider, error) {
	res, err := prahariTelemetry.NewResource(ctx, prahariTelemetry.Config{
		ServiceName: "integration-service",
		Environment: env,
		Version:     "1.0.0",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to build OTel resource: %w", err)
	}
	stdoutExp := prahariExporters.NewStdoutExporter()
	tp, err := prahariTelemetry.InitTracerProvider(res, prahariTelemetry.GetSampler(1.0), stdoutExp)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize tracer provider: %w", err)
	}
	return tp, nil
}

func initDatabase(ctx context.Context, databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	_ = db.PingContext(ctx)
	return db, nil
}

func initRouter(handler http.Handler) http.Handler {
	opts := prahariCORS.DefaultOptions()
	pipeline := prahariMid.New(
		prahariRequest.Middleware,
		prahariCORS.Middleware(opts),
		prahariLogMid.Middleware,
		prahariRecovery.Middleware,
		prahariTimeout.Middleware(15*time.Second),
		httpIface.PlantIsolationMiddleware,
	)
	return pipeline.Then(handler)
}

func RunApp() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := LoadConfig(ctx)
	if err != nil {
		panic(err)
	}

	tp, err := initTelemetry(ctx, cfg.Environment)
	if err == nil && tp != nil {
		defer func() { _ = tp.Shutdown(ctx) }()
	}

	db, err := initDatabase(ctx, cfg.DatabaseURL)
	if err != nil {
		prahariLogger.Error(ctx, "Database initialization failed", prahariLogger.Err(err))
	}
	if db != nil {
		defer db.Close()
	}

	redisClient, err := prahariRedis.NewClient(prahariRedis.Config{Address: cfg.RedisAddr})
	if err == nil && redisClient != nil {
		defer redisClient.Close()
	}

	store := pgInfra.NewStore(db)
	_ = redisInfra.NewCache(redisClient)
	kafkaPublisher := kafkaInfra.NewPublisher()

	// DI wiring
	connectorsSvc := connectorsApp.NewService(store, kafkaPublisher)
	syncSvc := synchronizationApp.NewService(store, kafkaPublisher)
	transformSvc := transformationApp.NewService(store, kafkaPublisher)
	reportingSvc := reportingApp.NewService(store)
	analyticsSvc := analyticsApp.NewService(store)
	searchSvc := searchApp.NewService(store)
	exportSvc := exportApp.NewService(store)

	schedulerSvc := schedulingApp.NewService()
	schedulerSvc.StartScheduler(ctx)

	// Event listener
	kafkaConsumer := kafkaInfra.NewConsumer()
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	// Outbound clients mapping E.g., MQTT routing
	_ = mqttInfra.NewClient(cfg.MqttBroker)

	// gRPC Server setup
	_ = grpcIface.NewServer(connectorsSvc, syncSvc, transformSvc, searchSvc)

	// HTTP
	mux := http.NewServeMux()
	handler := httpIface.NewHandler(connectorsSvc, syncSvc, reportingSvc, analyticsSvc, searchSvc, exportSvc)
	httpIface.RegisterRoutes(mux, handler)

	pipelineHandler := initRouter(mux)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      pipelineHandler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() { _ = srv.ListenAndServe() }()

	prahariLogger.Info(ctx, "Enterprise Integration Hub Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down Integration Service...")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()
	_ = srv.Shutdown(shutdownCtx)
	prahariLogger.Info(ctx, "Integration Service exited gracefully.")
}
