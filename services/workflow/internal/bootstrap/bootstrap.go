package bootstrap

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	prahariLogger "prahari/shared/logger"
	prahariApproval "prahari/services/workflow/internal/application/approval"
	prahariHTTP "prahari/services/workflow/internal/interfaces/http"
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
	tp, err := InitTelemetry(ctx, "workflow-service", cfg.Environment)
	if err == nil && tp != nil {
		defer func() {
			_ = tp.Shutdown(ctx)
		}()
	}

	// 4. Database Setup & Migrations
	_ = InitDatabase(ctx, cfg.DatabaseURL)

	// 5. Caching Setup
	redisClient, err := InitRedis(ctx, cfg.RedisAddr)
	if err == nil && redisClient != nil {
		defer redisClient.Close()
	}

	// 6. Messaging broker setups
	_ = InitKafka(ctx)

	// 7. Timer Scheduler setups
	InitScheduler(ctx)

	// 8. Dependency Injection Constructors Wiring
	approvalSvc := prahariApproval.NewService()
	httpHandler := prahariHTTP.NewHandler(approvalSvc)

	// 9. HTTP Routing Mux Setup
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks/complete", httpHandler.CompleteTask)

	// Chain middlewares pipeline
	pipelineHandler := InitRouter(mux)

	// 10. HTTP Server startup
	srv := StartServer(cfg.Port, pipelineHandler)

	prahariLogger.Info(ctx, "Workflow service initialized successfully. Server running.")

	// 11. Intercept SIGTERM / SIGINT for Graceful Shutdown
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
