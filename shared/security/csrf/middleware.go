package csrf

import (
	"crypto/subtle"
	"net/http"
)

const (
	CookieName = "csrf_token"
	HeaderName = "X-CSRF-Token"
)

// Middleware enforces CSRF checks via double-submit cookie patterns.
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Retrieve current cookie value
		cookie, err := r.Cookie(CookieName)
		var cookieToken string
		if err == nil {
			cookieToken = cookie.Value
		}

		// 2. For state-changing methods, validate the header matches cookie token
		if r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodDelete || r.Method == http.MethodPatch {
			headerToken := r.Header.Get(HeaderName)
			
			if cookieToken == "" || headerToken == "" {
				http.Error(w, "Forbidden: CSRF token missing", http.StatusForbidden)
				return
			}

			// Use subtle.ConstantTimeCompare to prevent timing side-channel leaks
			if subtle.ConstantTimeCompare([]byte(cookieToken), []byte(headerToken)) != 1 {
				http.Error(w, "Forbidden: CSRF token mismatch", http.StatusForbidden)
				return
			}
		}

		// 3. If cookie is missing, generate and set a new CSRF token for subsequent requests
		if cookieToken == "" {
			newToken, err := GenerateToken()
			if err == nil {
				http.SetCookie(w, &http.Cookie{
					Name:     CookieName,
					Value:    newToken,
					Path:     "/",
					HttpOnly: true,
					Secure:   r.URL.Scheme == "https" || r.TLS != nil,
					SameSite: http.SameSiteLaxMode,
					MaxAge:   86400, // 24 hours
				})
			}
		}

		next.ServeHTTP(w, r)
	})
}
