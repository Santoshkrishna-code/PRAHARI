package bootstrap

import (
	"context"
	"fmt"

	prahariTelemetry "prahari/shared/telemetry"
	prahariExporters "prahari/shared/telemetry/exporters"
	"go.opentelemetry.io/otel/sdk/trace"
)

// InitTelemetry sets up OTel resource markers and registers trace Providers.
func InitTelemetry(ctx context.Context, serviceName, env string) (*trace.TracerProvider, error) {
	res, err := prahariTelemetry.NewResource(ctx, prahariTelemetry.Config{
		ServiceName: serviceName,
		Environment: env,
		Version:     "1.0.0",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to build resource metadata: %w", err)
	}

	stdoutExp := prahariExporters.NewStdoutExporter()
	tp, err := prahariTelemetry.InitTracerProvider(res, prahariTelemetry.GetSampler(1.0), stdoutExp)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize global tracer: %w", err)
	}

	return tp, nil
}
