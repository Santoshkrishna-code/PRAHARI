package tracer

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// Tracer wraps OTel trace spans operations.
type Tracer struct {
	t trace.Tracer
}

// NewTracer constructs a new Tracer.
func NewTracer(name string) *Tracer {
	return &Tracer{
		t: otel.Tracer(name),
	}
}

// Start opens a new trace span linked to context.
func (tr *Tracer) Start(ctx context.Context, spanName string) (context.Context, trace.Span) {
	return tr.t.Start(ctx, spanName)
}
