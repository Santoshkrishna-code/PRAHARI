package middleware

import "net/http"

// SecurityHeaders injects standard web-hardening security headers.
func SecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Prevent Clickjacking attacks
		w.Header().Set("X-Frame-Options", "DENY")
		
		// Prevent MIME-sniffing
		w.Header().Set("X-Content-Type-Options", "nosniff")
		
		// Enforce HTTPS
		w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		
		// Control Referrer leaks
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		
		// XSS protection policy for older browsers
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		
		// Basic API Content Security Policy
		w.Header().Set("Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'; sandbox")
		
		next.ServeHTTP(w, r)
	})
}
