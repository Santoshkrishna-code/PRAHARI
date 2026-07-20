package headers

import (
	"net/http"
)

// MiddlewareOptions configures customizations for security headers.
type MiddlewareOptions struct {
	CSP               string
	HSTS              string
	PermissionsPolicy string
}

// DefaultOptions returns standard hardened configurations.
func DefaultOptions() MiddlewareOptions {
	csp := NewCSPBuilder().
		Add("default-src", "'none'").
		Add("frame-ancestors", "'none'").
		Add("sandbox").
		Build()

	hsts := HSTSConfig{
		MaxAge:            31536000, // 1 year
		IncludeSubDomains: true,
		Preload:           true,
	}.Build()

	permissions := NewPermissionsPolicyBuilder().
		Add("camera", "'none'").
		Add("microphone", "'none'").
		Add("geolocation", "'none'").
		Build()

	return MiddlewareOptions{
		CSP:               csp,
		HSTS:              hsts,
		PermissionsPolicy: permissions,
	}
}

// Middleware creates a handler that enforces all safety headers.
func Middleware(opts MiddlewareOptions) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Clickjacking protection
			w.Header().Set("X-Frame-Options", "DENY")

			// MIME-sniffing protection
			w.Header().Set("X-Content-Type-Options", "nosniff")

			// HSTS Enforces HTTPS
			if opts.HSTS != "" {
				w.Header().Set("Strict-Transport-Security", opts.HSTS)
			}

			// Referrer leaks prevention
			w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")

			// XSS Filter
			w.Header().Set("X-XSS-Protection", "1; mode=block")

			// CSP Content Security
			if opts.CSP != "" {
				w.Header().Set("Content-Security-Policy", opts.CSP)
			}

			// Hardware Access Permissions
			if opts.PermissionsPolicy != "" {
				w.Header().Set("Permissions-Policy", opts.PermissionsPolicy)
			}

			// Cross-Origin Isolations (COEP, COOP, CORP)
			w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
			w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
			w.Header().Set("Cross-Origin-Resource-Policy", "same-origin")

			next.ServeHTTP(w, r)
		})
	}
}
