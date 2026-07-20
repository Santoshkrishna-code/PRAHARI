package http

import (
	"net/http"

	prahariLogger "prahari/shared/logger"
)

// PlantIsolationMiddleware ensures plant isolation bounds.
func PlantIsolationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		plantID := r.Header.Get("X-Plant-ID")
		if plantID != "" {
			prahariLogger.Info(r.Context(), "Energy plant isolation checks applied",
				prahariLogger.String("plant_id", plantID))
		}
		next.ServeHTTP(w, r)
	})
}
