package auth

import (
	"encoding/json"
	"net/http"

	prahariJWT "prahari/shared/security/jwt"
	prahariMid "prahari/shared/middleware"
)

// Authenticator handles token validations and binds claims to HTTP requests contexts.
type Authenticator struct {
	validator *prahariJWT.Validator
}

// NewAuthenticator constructs an Authenticator.
func NewAuthenticator(v *prahariJWT.Validator) *Authenticator {
	return &Authenticator{validator: v}
}

// Middleware parses and validates incoming Bearer JWT tokens.
func (a *Authenticator) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr, err := prahariJWT.ExtractTokenFromHeader(r)
		if err != nil {
			a.respondUnauthorized(w, "Authentication credentials are required.")
			return
		}

		claims, err := a.validator.Validate(r.Context(), tokenStr)
		if err != nil {
			a.respondUnauthorized(w, "Invalid or expired authentication credentials.")
			return
		}

		// Inject claims into request context
		ctx := prahariMid.WithClaims(r.Context(), claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (a *Authenticator) respondUnauthorized(w http.ResponseWriter, detail string) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(http.StatusUnauthorized)
	
	problem := map[string]interface{}{
		"type":   "about:blank",
		"title":  "Unauthorized",
		"status": http.StatusUnauthorized,
		"detail": detail,
	}

	_ = json.NewEncoder(w).Encode(problem)
}
