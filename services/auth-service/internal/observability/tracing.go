package observability

import (
	"context"
	"fmt"
)

// TracingManager coordinates spans and traces lifecycle parameters.
type TracingManager struct {
	enabled bool
}

func NewTracingManager(enabled bool) *TracingManager {
	return &TracingManager{enabled: enabled}
}

// StartSpan returns context and finalizer to track execution flow.
func (t *TracingManager) StartSpan(ctx context.Context, name string) (context.Context, func()) {
	if !t.enabled {
		return ctx, func() {}
	}
	
	// In production, instantiate OpenTelemetry Tracer:
	// ctx, span := otel.Tracer("auth-service").Start(ctx, name)
	// return ctx, func() { span.End() }
	
	fmt.Printf("[TRACE-MOCK] Started execution span: %s\n", name)
	return ctx, func() {
		fmt.Printf("[TRACE-MOCK] Stopped execution span: %s\n", name)
	}
}
