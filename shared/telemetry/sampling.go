package telemetry

import (
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// GetSampler returns trace.Sampler conforming to the target ratio.
func GetSampler(ratio float64) sdktrace.Sampler {
	if ratio >= 1.0 {
		return sdktrace.AlwaysSample()
	}
	if ratio <= 0.0 {
		return sdktrace.NeverSample()
	}
	// Use ParentBased to honor upstream sampling choices while applying the local sampling ratio
	return sdktrace.ParentBased(sdktrace.TraceIDRatioBased(ratio))
}
