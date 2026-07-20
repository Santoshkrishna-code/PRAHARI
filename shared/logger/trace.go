package logger

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

// ExtractTraceContext inspects OpenTelemetry context spans and populates logging tags.
func ExtractTraceContext(ctx context.Context) context.Context {
	spanContext := trace.SpanContextFromContext(ctx)
	if spanContext.IsValid() {
		return WithTraceContext(
			ctx,
			spanContext.TraceID().String(),
			spanContext.SpanID().String(),
		)
	}
	return ctx
}
