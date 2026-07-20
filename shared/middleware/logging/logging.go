package logging

import (
	"net/http"
	"time"

	prahariLogger "prahari/shared/logger"
	prahariMid "prahari/shared/middleware"
)

type statusWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *statusWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// Middleware intercepting request latency and status codes.
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := &statusWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(sw, r)

		elapsed := time.Since(start)
		corrID := prahariMid.GetCorrelationID(r.Context())

		// Log request details using structured fields
		prahariLogger.Info(r.Context(), "HTTP Request Processed",
			prahariLogger.String("correlation_id", corrID),
			prahariLogger.String("method", r.Method),
			prahariLogger.String("path", r.URL.Path),
			prahariLogger.Int("status", sw.statusCode),
			prahariLogger.Duration("latency", elapsed),
		)
	})
}
