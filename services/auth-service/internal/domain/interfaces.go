package domain

import (
	"context"
)

// UserRepository represents the outbound adapter port for user identity management (Cognito).
type UserRepository interface {
	SignUp(ctx context.Context, email, password, role, firstName, lastName string) (string, error)
	SignIn(ctx context.Context, email, password string) (*TokenPair, error)
	RefreshToken(ctx context.Context, refreshToken string) (*TokenPair, error)
	GetUserByToken(ctx context.Context, accessToken string) (*User, error)
}

// TokenVerifier defines token verification operations.
type TokenVerifier interface {
	VerifyJWT(ctx context.Context, tokenStr string) (*JWTClaims, error)
}

// AuthUseCase represents the inbound application port for all authentication workflows.
type AuthUseCase interface {
	Register(ctx context.Context, email, password, role, firstName, lastName string) (*User, error)
	Login(ctx context.Context, email, password, ipAddress, userAgent string) (*TokenPair, error)
	Refresh(ctx context.Context, refreshToken string) (*TokenPair, error)
	GetUserProfile(ctx context.Context, accessToken string) (*User, error)
	VerifyToken(ctx context.Context, token string) (*JWTClaims, error)
}
