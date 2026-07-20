package metrics

import (
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

// Metrics wraps OTel meter instruments.
type Metrics struct {
	meter metric.Meter
}

// NewMetrics constructs a Metrics factory.
func NewMetrics(name string) *Metrics {
	return &Metrics{
		meter: otel.Meter(name),
	}
}

// NewCounter registers an atomic Int64 Counter (e.g. total request count).
func (m *Metrics) NewCounter(name, description string) (metric.Int64Counter, error) {
	counter, err := m.meter.Int64Counter(name, metric.WithDescription(description))
	if err != nil {
		return nil, fmt.Errorf("failed to register counter: %w", err)
	}
	return counter, nil
}

// NewHistogram registers a Float64 Histogram (e.g. request duration latency).
func (m *Metrics) NewHistogram(name, description string) (metric.Float64Histogram, error) {
	histogram, err := m.meter.Float64Histogram(name, metric.WithDescription(description))
	if err != nil {
		return nil, fmt.Errorf("failed to register histogram: %w", err)
	}
	return histogram, nil
}
