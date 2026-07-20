package middleware

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type contextKey string

const TraceparentKey contextKey = "traceparent"

// ExtractTracing resolves standard trace parent identifiers from message headers.
func ExtractTracing(ctx context.Context, headers []kafka.Header) context.Context {
	for _, h := range headers {
		if h.Key == "traceparent" {
			// Propagate traceparent value in context for child tracing
			return context.WithValue(ctx, TraceparentKey, string(h.Value))
		}
	}
	return ctx
}
