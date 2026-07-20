package telemetry

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

// InitTracerProvider sets up and registers the global TracerProvider.
func InitTracerProvider(res *resource.Resource, sampler trace.Sampler, exporter trace.SpanExporter) (*trace.TracerProvider, error) {
	tp := trace.NewTracerProvider(
		trace.WithSampler(sampler),
		trace.WithResource(res),
		trace.WithBatcher(exporter), // Batch spans in the background for performance
	)

	otel.SetTracerProvider(tp)
	return tp, nil
}

// InitMeterProvider sets up and registers the global MeterProvider.
func InitMeterProvider(res *resource.Resource, reader metric.Reader) (*metric.MeterProvider, error) {
	mp := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(reader),
	)

	otel.SetMeterProvider(mp)
	return mp, nil
}
