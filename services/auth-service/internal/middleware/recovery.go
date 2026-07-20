package middleware

import (
	"net/http"
	"runtime/debug"

	"go.uber.org/zap"
	"prahari/services/auth-service/internal/api"
)

// PanicRecovery handles panics inside handlers and responds with standard HTTP 500 errors.
func PanicRecovery(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					// Log the panic with stack trace
					logger.Error("Panic Recovered",
						zap.Any("error", err),
						zap.ByteString("stacktrace", debug.Stack()),
					)
					
					// Return clean error
					api.WriteError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "An unexpected server error occurred", nil)
				}
			}()
			
			next.ServeHTTP(w, r)
		})
	}
}
