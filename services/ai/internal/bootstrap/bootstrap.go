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
	analyticsApp "prahari/services/ai/internal/application/analytics"
	chatApp "prahari/services/ai/internal/application/chat"
	exportApp "prahari/services/ai/internal/application/export"
	predictionApp "prahari/services/ai/internal/application/prediction"
	ragApp "prahari/services/ai/internal/application/rag"
	recommendationApp "prahari/services/ai/internal/application/recommendation"
	searchApp "prahari/services/ai/internal/application/search"
	summarizationApp "prahari/services/ai/internal/application/summarization"

	// Infrastructure adapters
	embeddingsInfra "prahari/services/ai/internal/infrastructure/embeddings"
	kafkaInfra "prahari/services/ai/internal/infrastructure/kafka"
	llmInfra "prahari/services/ai/internal/infrastructure/llm"
	pgInfra "prahari/services/ai/internal/infrastructure/postgres"
	redisInfra "prahari/services/ai/internal/infrastructure/redis"
	vectorInfra "prahari/services/ai/internal/infrastructure/vectorstore"

	// Interfaces
	eventsIface "prahari/services/ai/internal/interfaces/events"
	grpcIface "prahari/services/ai/internal/interfaces/grpc"
	httpIface "prahari/services/ai/internal/interfaces/http"
)

func initTelemetry(ctx context.Context, env string) (*trace.TracerProvider, error) {
	res, err := prahariTelemetry.NewResource(ctx, prahariTelemetry.Config{
		ServiceName: "ai-service",
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

	// Infrastructure wrappers
	vectorClient := vectorInfra.NewClient(cfg.VectorAddr)
	llmClient := llmInfra.NewClient(cfg.LLMModel)
	embedClient := embeddingsInfra.NewClient(cfg.EmbedModel)

	// DI wiring
	ragSvc := ragApp.NewService(store, vectorClient, embedClient, kafkaPublisher)
	chatSvc := chatApp.NewService(store, llmClient)
	recommendationSvc := recommendationApp.NewService(store, kafkaPublisher)
	summarizationSvc := summarizationApp.NewService(store, kafkaPublisher)
	predictionSvc := predictionApp.NewService(store, kafkaPublisher)
	analyticsSvc := analyticsApp.NewService()
	searchSvc := searchApp.NewService(store)
	exportSvc := exportApp.NewService(store)

	// Event listener
	kafkaConsumer := kafkaInfra.NewConsumer()
	eventListener := eventsIface.NewListener(kafkaConsumer)
	eventListener.Start(ctx)

	// gRPC Server setup
	_ = grpcIface.NewServer(chatSvc, ragSvc, summarizationSvc, recommendationSvc, predictionSvc, searchSvc)

	// HTTP
	mux := http.NewServeMux()
	handler := httpIface.NewHandler(ragSvc, chatSvc, recommendationSvc, summarizationSvc, predictionSvc, analyticsSvc, searchSvc, exportSvc)
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

	prahariLogger.Info(ctx, "AI Platform & Copilot Service initialized successfully",
		prahariLogger.Int("port", cfg.Port),
		prahariLogger.String("environment", cfg.Environment))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down AI Service...")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer shutdownCancel()
	_ = srv.Shutdown(shutdownCtx)
	prahariLogger.Info(ctx, "AI Service exited gracefully.")
}
