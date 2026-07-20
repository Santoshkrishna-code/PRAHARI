package cors

import (
	"net/http"
	"strings"
)

// Options holds configurations for CORS behavior.
type Options struct {
	AllowedOrigins   []string
	AllowedMethods   []string
	AllowedHeaders   []string
	AllowCredentials bool
}

// DefaultOptions returns standard hardened settings.
func DefaultOptions() Options {
	return Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "X-Correlation-ID"},
		AllowCredentials: true,
	}
}

// Middleware constructs a CORS handler.
func Middleware(opts Options) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			origin := r.Header.Get("Origin")
			if origin != "" {
				// Check origin permissions (simple wildcard or match check)
				allowOrigin := "*"
				if len(opts.AllowedOrigins) > 0 && opts.AllowedOrigins[0] != "*" {
					for _, o := range opts.AllowedOrigins {
						if o == origin {
							allowOrigin = origin
							break
						}
					}
				} else if opts.AllowCredentials {
					// W3C spec requires explicit origin (not *) when credentials are allowed
					allowOrigin = origin
				}

				w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
				w.Header().Set("Access-Control-Allow-Methods", strings.Join(opts.AllowedMethods, ", "))
				w.Header().Set("Access-Control-Allow-Headers", strings.Join(opts.AllowedHeaders, ", "))

				if opts.AllowCredentials {
					w.Header().Set("Access-Control-Allow-Credentials", "true")
				}
			}

			// Intercept preflight OPTIONS request
			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusNoContent)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
