package grpc

import (
	"context"
	"errors"

	prahariJWT "prahari/shared/security/jwt"
)

// Server implements gRPC authentication verifications.
type Server struct {
	validator *prahariJWT.Validator
}

// NewServer constructs a Server.
func NewServer(v *prahariJWT.Validator) *Server {
	return &Server{validator: v}
}

// ValidateToken verifies token signatures, returning scope claims.
func (s *Server) ValidateToken(ctx context.Context, token string) (*prahariJWT.Claims, error) {
	if token == "" {
		return nil, errors.New("authentication token is required")
	}

	claims, err := s.validator.Validate(ctx, token)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
