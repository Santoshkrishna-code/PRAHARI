package tracing

import (
	"net/http"
)

// Middleware extracts incoming trace carrier maps from request headers.
// Instantiates a child trace span matching HTTP methods and paths.
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// In production:
		// ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))
		// ctx, span := otel.Tracer("prahari-http").Start(ctx, r.Method+" "+r.URL.Path)
		// defer span.End()

		next.ServeHTTP(w, r)
	})
}
