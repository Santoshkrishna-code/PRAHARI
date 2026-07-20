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
	analyticsApp "prahari/services/digitaltwin/internal/application/analytics"
	exportApp "prahari/services/digitaltwin/internal/application/export"
	playbackApp "prahari/services/digitaltwin/internal/application/playback"
	searchApp "prahari/services/digitaltwin/internal/application/search"
	simulationApp "prahari/services/digitaltwin/internal/application/simulation"
	syncApp "prahari/services/digitaltwin/internal/application/synchronization"
	visualizationApp "prahari/services/digitaltwin/internal/application/visualization"

	// Infrastructure adapters
	graphstoreInfra "prahari/services/digitaltwin/internal/infrastructure/graphstore"
	kafkaInfra "prahari/services/digitaltwin/internal/infrastructure/kafka"
	pgInfra "prahari/services/digitaltwin/internal/infrastructure/postgres"
	redisInfra "prahari/services/digitaltwin/internal/infrastructure/redis"
	websocketInfra "prahari/services/digitaltwin/internal/infrastructure/websocket"

	// Interfaces
	eventsIface "prahari/services/digitaltwin/internal/interfaces/events"
	grpcIface "prahari/services/digitaltwin/internal/interfaces/grpc"
	httpIface "prahari/services/digitaltwin/internal/interfaces/http"
)

func initTelemetry(ctx context.Context, env string) (*trace.TracerProvider, error) {
	res, err := prahariTelemetry.NewResource(ctx, prahariTelemetry.Config{
		ServiceName: "digitaltwin-service",
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
		httpIface.TenantIsolationMiddleware,
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
	syncSvc := syncApp.NewService(store, kafkaPublisher)
	visualizationSvc := visualizationApp.NewService(store)
	simulationSvc := simulationApp.NewService(store, kafkaPublisher)
	playbackSvc := playbackApp.NewService(store, kafkaPublisher)
	analyticsSvc := analyticsApp.NewService()
	searchSvc := searchApp.NewService(store)
	exportSvc := exportApp.NewService(store)

	// Event listener
	kafkaConsumer := kafkaInfra.NewConsumer()
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	// Outbound client connections
	_ = websocketInfra.NewHub()
	_ = graphstoreInfra.NewClient(cfg.GraphAddr)

	// gRPC Server setup
	_ = grpcIface.NewServer(syncSvc, simulationSvc, playbackSvc, searchSvc)

	// HTTP
	mux := http.NewServeMux()
	handler := httpIface.NewHandler(syncSvc, visualizationSvc, simulationSvc, playbackSvc, analyticsSvc, searchSvc, exportSvc)
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

	prahariLogger.Info(ctx, "Digital Twin Platform Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down Digital Twin Service...")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()
	_ = srv.Shutdown(shutdownCtx)
	prahariLogger.Info(ctx, "Digital Twin Service exited gracefully.")
}
