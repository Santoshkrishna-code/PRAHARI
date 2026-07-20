package requestid

import (
	"crypto/rand"
	"fmt"
	"net/http"

	prahariMid "prahari/shared/middleware"
)

const (
	HeaderName = "X-Correlation-ID"
)

// Middleware extracts the correlation ID header or generates a new one.
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		corrID := r.Header.Get(HeaderName)
		if corrID == "" {
			corrID = generateUUID()
		}

		// Inject into context
		ctx := prahariMid.WithCorrelationID(r.Context(), corrID)

		// Set response header to propagate ID back to the client
		w.Header().Set(HeaderName, corrID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func generateUUID() string {
	uuid := make([]byte, 16)
	_, _ = rand.Read(uuid)
	// Set version and variant bits for v4 compliance
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}
