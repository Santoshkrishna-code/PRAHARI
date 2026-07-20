package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	prahariConfig "prahari/shared/config"
	prahariLogger "prahari/shared/logger"
	prahariRedis "prahari/shared/redis"
	prahariMid "prahari/shared/middleware"
	prahariRequest "prahari/shared/middleware/requestid"
	prahariLogMid "prahari/shared/middleware/logging"
	prahariRecovery "prahari/shared/middleware/recovery"
	prahariTelemetry "prahari/shared/telemetry"
	prahariExporters "prahari/shared/telemetry/exporters"
	prahariTracer "prahari/shared/telemetry/tracer"
)

type AppConfig struct {
	Port        int    `env:"PORT" envDefault:"8080"`
	RedisAddr   string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
	Environment string `env:"ENVIRONMENT" envDefault:"development"`
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 1. Initialize Structured Logger
	prahariLogger.Info(ctx, "Starting microservice bootstrap...")

	// 2. Load Configuration
	loader, err := prahariConfig.NewLoader(prahariConfig.Options{})
	if err != nil {
		prahariLogger.Error(ctx, "Failed to initialize configuration loader", prahariLogger.Err(err))
		return
	}

	var cfg AppConfig
	if err := loader.Load(ctx, &cfg); err != nil {
		prahariLogger.Error(ctx, "Failed to load configuration overrides", prahariLogger.Err(err))
		return
	}
	prahariLogger.Info(ctx, fmt.Sprintf("Loaded settings. Listening on port: %d", cfg.Port))

	// 3. Initialize Redis Client
	redisClient, err := prahariRedis.NewClient(prahariRedis.Config{
		Address: cfg.RedisAddr,
	})
	if err != nil {
		prahariLogger.Error(ctx, "Failed to connect to Redis cache cluster", prahariLogger.Err(err))
		return
	}
	defer redisClient.Close()
	prahariLogger.Info(ctx, "Connected to Redis cluster.")

	// 4. Initialize Telemetry provider
	res, err := prahariTelemetry.NewResource(ctx, prahariTelemetry.Config{
		ServiceName: "example-bootstrap-service",
		Environment: cfg.Environment,
		Version:     "1.0.0",
	})
	if err != nil {
		prahariLogger.Error(ctx, "Failed to build telemetry resource attributes", prahariLogger.Err(err))
		return
	}

	stdoutExp := prahariExporters.NewStdoutExporter()
	tp, _ := prahariTelemetry.InitTracerProvider(res, prahariTelemetry.GetSampler(1.0), stdoutExp)
	defer func() {
		_ = tp.Shutdown(ctx)
	}()

	tr := prahariTracer.NewTracer("example-service")
	prahariLogger.Info(ctx, "Telemetry and Trace Providers registered.")

	// 5. Setup Middlewares & Routes
	pipeline := prahariMid.New(
		prahariRequest.Middleware,
		prahariLogMid.Middleware,
		prahariRecovery.Middleware,
	)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// Run within a child trace span
		reqCtx, span := tr.Start(r.Context(), "ServeHealth")
		defer span.End()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ONLINE"}`))

		prahariLogger.Info(reqCtx, "Health status query responded.")
	})

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      pipeline.Then(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	prahariLogger.Info(ctx, fmt.Sprintf("HTTP Server listening on address %s", server.Addr))
	// In execution environments, start listen loop:
	// _ = server.ListenAndServe()
}
