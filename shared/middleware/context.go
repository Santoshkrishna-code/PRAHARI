package middleware

import (
	"context"
)

type contextKey string

const (
	correlationIDKey contextKey = "prahari_correlation_id"
	claimsKey        contextKey = "prahari_jwt_claims"
)

// WithCorrelationID binds a correlation ID string to the context.
func WithCorrelationID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, correlationIDKey, id)
}

// GetCorrelationID extracts the correlation ID from the context.
func GetCorrelationID(ctx context.Context) string {
	if val, ok := ctx.Value(correlationIDKey).(string); ok {
		return val
	}
	return ""
}

// WithClaims binds JWT claims data to the context.
func WithClaims(ctx context.Context, claims interface{}) context.Context {
	return context.WithValue(ctx, claimsKey, claims)
}

// GetClaims extracts JWT claims from the context.
func GetClaims(ctx context.Context) interface{} {
	return ctx.Value(claimsKey)
}
