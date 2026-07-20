package recovery

import (
	"encoding/json"
	"fmt"
	"net/http"

	prahariLogger "prahari/shared/logger"
)

// Middleware intercepts runtime panics and converts them into standardized JSON error responses.
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				// Log the panic structure
				err := fmt.Errorf("panic: %v", rec)
				prahariLogger.Error(r.Context(), "runtime panic recovered in HTTP pipeline", prahariLogger.Err(err))

				// Construct clean RFC 7807 response
				w.Header().Set("Content-Type", "application/problem+json")
				w.WriteHeader(http.StatusInternalServerError)

				problem := map[string]interface{}{
					"type":   "about:blank",
					"title":  "Internal Server Error",
					"status": http.StatusInternalServerError,
					"detail": "An unexpected error occurred on the server.",
				}

				_ = json.NewEncoder(w).Encode(problem)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
