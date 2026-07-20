package logger

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

// LogRequest logs HTTP request parameters and latency metrics.
func LogRequest(log *Logger, r *http.Request, status, size int, latency time.Duration) {
	if log == nil {
		return
	}

	log.WithContext(r.Context()).Info("HTTP Request",
		zap.String("method", r.Method),
		zap.String("path", r.URL.Path),
		zap.String("query", r.URL.RawQuery),
		zap.String("remote_ip", r.RemoteAddr),
		zap.Int("status", status),
		zap.Int("response_size_bytes", size),
		zap.Duration("latency", latency),
	)
}
