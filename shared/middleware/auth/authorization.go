package auth

import (
	"encoding/json"
	"net/http"

	prahariMid "prahari/shared/middleware"
	prahariJWT "prahari/shared/security/jwt"
)

// AuthorizeRole restricts access to requests holding specific allowed roles.
func AuthorizeRole(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claimsVal := prahariMid.GetClaims(r.Context())
			if claimsVal == nil {
				respondForbidden(w, "Access denied: unauthenticated request.")
				return
			}

			claims, ok := claimsVal.(*prahariJWT.Claims)
			if !ok {
				respondForbidden(w, "Access denied: invalid token claims format.")
				return
			}

			// Check role permissions
			authorized := false
			for _, r := range allowedRoles {
				if r == claims.Role {
					authorized = true
					break
				}
			}

			if !authorized {
				respondForbidden(w, "Access denied: insufficient permissions.")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func respondForbidden(w http.ResponseWriter, detail string) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(http.StatusForbidden)

	problem := map[string]interface{}{
		"type":   "about:blank",
		"title":  "Forbidden",
		"status": http.StatusForbidden,
		"detail": detail,
	}

	_ = json.NewEncoder(w).Encode(problem)
}
