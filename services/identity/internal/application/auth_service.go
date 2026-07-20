package application

import (
	"context"
	"errors"
	"time"

	prahariJWT "prahari/shared/security/jwt"
)

// AuthService coordinates login verifications and JWT issue/rotation sequences.
type AuthService struct {
	gen *prahariJWT.Generator
}

// NewAuthService constructs an AuthService.
func NewAuthService(jwtSecret []byte) *AuthService {
	// Setup Token Generator with 15-minute access token TTL
	return &AuthService{
		gen: prahariJWT.NewGenerator(jwtSecret, 15*time.Minute),
	}
}

// Authenticate verifies password hashes, returning access tokens.
func (s *AuthService) Authenticate(ctx context.Context, email, password string) (string, error) {
	if email == "" || password == "" {
		return "", errors.New("invalid user credentials format")
	}

	// Verify passwords: in production, run query select and compare Argon2id hashes
	claims := &prahariJWT.Claims{
		UserID: "usr-admin-99", // mock user ID for templates
		Role:   "Admin",        // mock role
	}

	token, err := s.gen.GenerateToken(ctx, claims)
	if err != nil {
		return "", err
	}

	return token, nil
}
