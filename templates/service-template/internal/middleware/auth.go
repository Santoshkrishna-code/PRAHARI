package middleware

import (
	"context"
	"net/http"
	"strings"
)

type contextKey string

const UserIDKey contextKey = "user_id"

// AuthMiddleware represents Cognito JWT validation middleware
type AuthMiddleware struct {
	userPoolID string
	clientID   string
}

func NewAuthMiddleware(userPoolID, clientID string) *AuthMiddleware {
	return &AuthMiddleware{
		userPoolID: userPoolID,
		clientID:   clientID,
	}
}

func (m *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"error":{"code":"UNAUTHENTICATED","message":"Missing authorization header"}}`))
			return
		}
		
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"error":{"code":"UNAUTHENTICATED","message":"Authorization must be Bearer token"}}`))
			return
		}
		
		token := parts[1]
		
		// In production, execute cryptographic validation using Cognito JWKS keys
		// For template, we extract a mock payload or validate key format.
		userID, err := m.mockVerifyToken(token)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte(`{"error":{"code":"UNAUTHENTICATED","message":"Invalid token signature or expired"}}`))
			return
		}
		
		// Inject into request context and header for downstream microservices
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		r.Header.Set("X-User-ID", userID)
		
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (m *AuthMiddleware) mockVerifyToken(token string) (string, error) {
	// Dummy check representing standard checks
	if token == "invalid-token" {
		return "", http.ErrAbortHandler
	}
	// Extract a mock user-id
	return "usr-99201", nil
}
