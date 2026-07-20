package domain

import (
	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims represents parsed Cognito access/id token payload metadata.
type JWTClaims struct {
	jwt.RegisteredClaims
	Email    string   `json:"email"`
	Role     UserRole `json:"custom:role"` // Assumes Cognito custom attribute mapping or group list
	Groups   []string `json:"cognito:groups"`
	Username string   `json:"username"`
}
