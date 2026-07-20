package logger_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
	prahariLogger "prahari/shared/logger"
)

func TestNewLogger_Development(t *testing.T) {
	opts := prahariLogger.DefaultOptions()
	opts.ServiceName = "test-service"
	
	log, err := prahariLogger.New(opts)
	if err != nil {
		t.Fatalf("failed to instantiate logger: %v", err)
	}
	
	if log.GetZapLogger() == nil {
		t.Fatal("underlying Zap logger was nil")
	}
}

func TestLogger_WithContext_CorrelationID(t *testing.T) {
	// Create an observer core to capture logged entries in memory
	core, logs := observer.New(zap.DebugLevel)
	zapLogger := zap.New(core)
	
	opts := prahariLogger.DefaultOptions()
	log := &prahariLogger.Logger{
		Logger: zapLogger,
	}

	// Inject Correlation ID into context
	ctx := context.Background()
	ctx = prahariLogger.WithCorrelationID(ctx, "corr-12345")

	// Log message using context
	log.WithContext(ctx).Info("test context log")

	// Validate log capture contains field
	entries := logs.All()
	if len(entries) != 1 {
		t.Fatalf("expected 1 log entry, got %d", len(entries))
	}

	correlationIDFound := false
	for _, field := range entries[0].Context {
		if field.Key == "correlation_id" && field.String == "corr-12345" {
			correlationIDFound = true
			break
		}
	}

	if !correlationIDFound {
		t.Error("expected logged metadata field correlation_id to be populated")
	}
}

func TestLogRequest(t *testing.T) {
	core, logs := observer.New(zap.InfoLevel)
	zapLogger := zap.New(core)
	
	log := &prahariLogger.Logger{
		Logger: zapLogger,
	}

	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login?test=1", nil)
	prahariLogger.LogRequest(log, req, http.StatusOK, 1024, 5*time.Millisecond)

	entries := logs.All()
	if len(entries) != 1 {
		t.Fatalf("expected 1 log entry, got %d", len(entries))
	}

	entry := entries[0]
	if entry.Message != "HTTP Request" {
		t.Errorf("expected message 'HTTP Request', got '%s'", entry.Message)
	}
}
