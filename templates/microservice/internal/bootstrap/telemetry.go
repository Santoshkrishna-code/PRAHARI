package bootstrap

import (
	"context"
	"fmt"

	prahariTelemetry "prahari/shared/telemetry"
	prahariExporters "prahari/shared/telemetry/exporters"
	"go.opentelemetry.io/otel/sdk/trace"
)

// InitTelemetry registers resource attributes and trace providers.
func InitTelemetry(ctx context.Context, serviceName, env string) (*trace.TracerProvider, error) {
	res, err := prahariTelemetry.NewResource(ctx, prahariTelemetry.Config{
		ServiceName: serviceName,
		Environment: env,
		Version:     "1.0.0",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	// Default to console exporter in templates. Swap with OTLP exporters in productions.
	stdoutExp := prahariExporters.NewStdoutExporter()
	tp, err := prahariTelemetry.InitTracerProvider(res, prahariTelemetry.GetSampler(1.0), stdoutExp)
	if err != nil {
		return nil, fmt.Errorf("failed to construct trace provider: %w", err)
	}

	return tp, nil
}
