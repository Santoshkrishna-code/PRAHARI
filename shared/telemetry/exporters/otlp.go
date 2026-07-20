package exporters

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/trace"
)

// NewOTLPExporter instantiates an OTLP gRPC exporter.
func NewOTLPExporter(ctx context.Context, endpoint string) (trace.SpanExporter, error) {
	// For production configurations, replace WithInsecure with secure TLS cert options
	exporter, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(endpoint),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to initialize OTLP gRPC exporter: %w", err)
	}

	return exporter, nil
}
