package middleware

import (
	"context"
	"net/http"
	"strings"

	"prahari/services/auth-service/internal/api"
	"prahari/services/auth-service/internal/domain"
	"prahari/services/auth-service/internal/utils"
)

type contextKey string

const (
	UserIDKey contextKey = "user_id"
	EmailKey  contextKey = "email"
	RoleKey   contextKey = "role"
)

// AuthenticateToken verifies the Bearer JWT and propagates claims in request contexts.
func AuthenticateToken(verifier *utils.TokenVerifier) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				api.WriteError(w, http.StatusUnauthorized, "UNAUTHENTICATED", "Missing authorization header", nil)
				return
			}
			
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				api.WriteError(w, http.StatusUnauthorized, "UNAUTHENTICATED", "Authorization header must be Bearer token", nil)
				return
			}
			
			tokenStr := parts[1]
			claims, err := verifier.Verify(r.Context(), tokenStr)
			if err != nil {
				api.MapError(w, err)
				return
			}
			
			// Inject claims into context
			ctx := context.WithValue(r.Context(), UserIDKey, claims.Subject)
			ctx = context.WithValue(ctx, EmailKey, claims.Email)
			ctx = context.WithValue(ctx, RoleKey, string(claims.Role))
			
			// Add context headers for upstream proxy validation
			r.Header.Set("X-User-ID", claims.Subject)
			r.Header.Set("X-User-Email", claims.Email)
			r.Header.Set("X-User-Role", string(claims.Role))
			
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// GetUserFromContext extracts active user metadata fields from contexts.
func GetUserFromContext(ctx context.Context) (*domain.User, bool) {
	id, ok1 := ctx.Value(UserIDKey).(string)
	email, ok2 := ctx.Value(EmailKey).(string)
	role, ok3 := ctx.Value(RoleKey).(string)
	
	if !ok1 || !ok2 || !ok3 {
		return nil, false
	}
	
	return &domain.User{
		ID:    id,
		Email: email,
		Role:  domain.UserRole(role),
	}, true
}
