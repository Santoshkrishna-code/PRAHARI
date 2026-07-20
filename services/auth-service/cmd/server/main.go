package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"prahari/services/auth-service/bootstrap"
)

func main() {
	// Root context for application lifecycle
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 1. Build Dependency Injection Container
	app, err := bootstrap.BuildApplication(ctx)
	if err != nil {
		// Fallback logger if bootstrap fails
		println("FATAL: Failed to bootstrap application: " + err.Error())
		os.Exit(1)
	}

	app.Logger.Info("Bootstrap complete. Launching server...")

	// 2. Configure HTTP Server
	server := &http.Server{
		Addr:         ":" + app.Config.Server.HTTPPort,
		Handler:      app.Router,
		ReadTimeout:  app.Config.Server.ReadTimeout,
		WriteTimeout: app.Config.Server.WriteTimeout,
		IdleTimeout:  app.Config.Server.IdleTimeout,
	}

	// 3. Start Server asynchronously to allow graceful shutdown handling
	go func() {
		app.Logger.Info("HTTP Server listening", zap.String("port", app.Config.Server.HTTPPort))
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			app.Logger.Fatal("HTTP Server failed", zap.Error(err))
		}
	}()

	// 4. Wait for termination signals
	<-ctx.Done()
	app.Logger.Info("Shutdown signal received. Initiating graceful shutdown...")

	// 5. Execute Graceful Shutdown with timeout context
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		app.Logger.Error("Server forced to shutdown", zap.Error(err))
	} else {
		app.Logger.Info("Server stopped cleanly.")
	}
}
