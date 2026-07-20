package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

// responseWriterWrapper tracks HTTP status code and response size metrics.
type responseWriterWrapper struct {
	http.ResponseWriter
	status int
	size   int
}

func (w *responseWriterWrapper) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseWriterWrapper) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = http.StatusOK
	}
	n, err := w.ResponseWriter.Write(b)
	w.size += n
	return n, err
}

// StructuredLogging returns a middleware that logs incoming requests with Zap.
func StructuredLogging(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			
			wrapper := &responseWriterWrapper{ResponseWriter: w}
			next.ServeHTTP(wrapper, r)
			
			latency := time.Since(start)
			
			logger.Info("HTTP Request",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.String("remote_ip", r.RemoteAddr),
				zap.Int("status", wrapper.status),
				zap.Int("size_bytes", wrapper.size),
				zap.Duration("latency", latency),
			)
		})
	}
}
