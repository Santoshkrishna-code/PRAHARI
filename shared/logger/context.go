package logger

import (
	"context"

	"go.uber.org/zap"
)

type contextKey string

const (
	CorrelationIDKey contextKey = "correlation_id"
	TraceIDKey       contextKey = "trace_id"
	SpanIDKey        contextKey = "span_id"
)

// WithCorrelationID binds a correlation ID string into context.
func WithCorrelationID(ctx context.Context, correlationID string) context.Context {
	return context.WithValue(ctx, CorrelationIDKey, correlationID)
}

// GetCorrelationID retrieves the active correlation ID.
func GetCorrelationID(ctx context.Context) string {
	if val, ok := ctx.Value(CorrelationIDKey).(string); ok {
		return val
	}
	return ""
}

// WithTraceContext binds OpenTelemetry trace and span contexts.
func WithTraceContext(ctx context.Context, traceID, spanID string) context.Context {
	ctx = context.WithValue(ctx, TraceIDKey, traceID)
	return context.WithValue(ctx, SpanIDKey, spanID)
}

// FromContext extracts contextual identifiers and returns a child logger.
func FromContext(ctx context.Context, baseLogger *zap.Logger) *zap.Logger {
	if ctx == nil || baseLogger == nil {
		return baseLogger
	}
	
	fields := make([]zap.Field, 0, 3)
	
	if corrID := GetCorrelationID(ctx); corrID != "" {
		fields = append(fields, zap.String("correlation_id", corrID))
	}
	
	if traceID, ok := ctx.Value(TraceIDKey).(string); ok && traceID != "" {
		fields = append(fields, zap.String("trace_id", traceID))
	}
	
	if spanID, ok := ctx.Value(SpanIDKey).(string); ok && spanID != "" {
		fields = append(fields, zap.String("span_id", spanID))
	}
	
	if len(fields) > 0 {
		return baseLogger.With(fields...)
	}
	
	return baseLogger
}
