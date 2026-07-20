package metrics_test

import (
	"context"
	"testing"

	"prahari/shared/telemetry/metrics"
)

func TestMetricsRegistry(t *testing.T) {
	m := metrics.NewMetrics("test-metrics-service")
	ctx := context.Background()

	counter, err := m.NewCounter("http_requests_total", "total requests counter")
	if err != nil {
		t.Fatalf("failed to create counter: %v", err)
	}

	if counter == nil {
		t.Fatal("expected counter to be instantiated, got nil")
	}

	// Increment metric
	counter.Add(ctx, 1)
}
