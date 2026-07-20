package http

import "net/http"

func PlantIsolationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		plantID := r.Header.Get("X-Plant-ID")
		if plantID != "" {
			r.Header.Set("X-Isolated-Plant-ID", plantID)
		}
		next.ServeHTTP(w, r)
	})
}
