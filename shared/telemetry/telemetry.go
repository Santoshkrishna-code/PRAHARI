package telemetry

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Manager aggregates providers and coordinates graceful shutdowns.
type Manager struct {
	tp *trace.TracerProvider
	mp *metric.MeterProvider
}

// NewManager constructs a telemetry manager.
func NewManager(tp *trace.TracerProvider, mp *metric.MeterProvider) *Manager {
	return &Manager{
		tp: tp,
		mp: mp,
	}
}

// Shutdown flushes all queued traces and metrics streams before container exits.
func (m *Manager) Shutdown(ctx context.Context) error {
	var errs []error

	if m.tp != nil {
		if err := m.tp.Shutdown(ctx); err != nil {
			errs = append(errs, fmt.Errorf("failed to shutdown tracer provider: %w", err))
		}
	}

	if m.mp != nil {
		if err := m.mp.Shutdown(ctx); err != nil {
			errs = append(errs, fmt.Errorf("failed to shutdown meter provider: %w", err))
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("telemetry shutdown encountered errors: %v", errs)
	}

	return nil
}
