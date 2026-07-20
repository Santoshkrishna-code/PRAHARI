package bootstrap

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	prahariLogger "prahari/shared/logger"
	prahariRouting "prahari/edge/gateway/internal/application/routing"
	prahariHTTP "prahari/edge/gateway/internal/interfaces/http"
)

// RunApp runs configurations overrides, initializes connections, and starts HTTP.
func RunApp() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 1. Config Loader
	cfg, err := LoadConfig(ctx)
	if err != nil {
		panic(err)
	}

	// 2. Logger Setup
	_ = InitLogger(cfg.Environment)

	// 3. Telemetry Register
	tp, err := InitTelemetry(ctx, "gateway-service", cfg.Environment)
	if err == nil && tp != nil {
		defer func() {
			_ = tp.Shutdown(ctx)
		}()
	}

	// 4. Caching Setup
	redisClient, err := InitRedis(ctx, cfg.RedisAddr)
	if err == nil && redisClient != nil {
		defer redisClient.Close()
	}

	// 5. Dependency Injection Constructors Wiring
	routeTable := prahariRouting.NewTable(cfg.Routes)
	gatewayHandler := prahariHTTP.NewHandler(routeTable)

	// 6. HTTP Routing Mux Setup
	mux := http.NewServeMux()
	mux.Handle("/", gatewayHandler)

	// Chain middlewares pipeline
	pipelineHandler := InitRouter(mux)

	// 7. HTTP Server startup
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "Gateway service initialized successfully. Server running.")

	// 8. Intercept SIGTERM / SIGINT for Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	prahariLogger.Info(ctx, "Shutting down HTTP server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		prahariLogger.Error(ctx, "HTTP server shutdown failed", prahariLogger.Err(err))
	}

	prahariLogger.Info(ctx, "Server exited gracefully.")
}
