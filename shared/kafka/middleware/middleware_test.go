package middleware_test

import (
	"context"
	"errors"
	"testing"

	"github.com/segmentio/kafka-go"
	"prahari/shared/kafka/middleware"
)

func TestLoggingAndMetricsMiddlewares(t *testing.T) {
	tracker := &middleware.MetricsTracker{}
	
	// Create handler chain
	handler := middleware.LoggingMiddleware(func(ctx context.Context, key, val []byte) error {
		if string(key) == "fail" {
			return errors.New("processing error")
		}
		return nil
	})

	metricsHandler := tracker.MetricsMiddleware(handler)

	ctx := context.Background()

	// Case 1: Succeeded
	err := metricsHandler(ctx, []byte("ok"), []byte("payload"))
	if err != nil {
		t.Fatalf("expected handler to succeed, got: %v", err)
	}

	if tracker.SuccessCount != 1 {
		t.Errorf("expected 1 success count, got %d", tracker.SuccessCount)
	}

	// Case 2: Failed
	err = metricsHandler(ctx, []byte("fail"), []byte("payload"))
	if err == nil {
		t.Fatal("expected handler to return error, got nil")
	}

	if tracker.FailureCount != 1 {
		t.Errorf("expected 1 failure count, got %d", tracker.FailureCount)
	}
}

func TestTracing_Extract(t *testing.T) {
	headers := []kafka.Header{
		{Key: "traceparent", Value: []byte("00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01")},
	}

	ctx := middleware.ExtractTracing(context.Background(), headers)
	val := ctx.Value(middleware.TraceparentKey)
	if val == nil {
		t.Fatal("expected traceparent header to be extracted into context, got nil")
	}

	if val.(string) != "00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01" {
		t.Errorf("extracted traceparent value mismatch, got: %s", val.(string))
	}
}
