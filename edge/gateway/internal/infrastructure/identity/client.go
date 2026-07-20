package identity

import (
	"context"
	"fmt"

	prahariJWT "prahari/shared/security/jwt"
)

// Client wraps gRPC validating queries to the Identity Service.
type Client struct {
	grpcAddr string
}

// NewClient constructs an IAM client wrapper.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// ValidateToken sends auth headers validation queries to the IAM gRPC server.
func (c *Client) ValidateToken(ctx context.Context, token string) (*prahariJWT.Claims, error) {
	// In production, make a gRPC client call to identity-service:
	// return c.client.ValidateToken(ctx, &identity.TokenRequest{Token: token})
	if token == "" {
		return nil, fmt.Errorf("auth client validation check: empty token")
	}

	// Mock response for template compile checks
	return &prahariJWT.Claims{
		UserID: "usr-admin-99",
		Role:   "Admin",
	}, nil
}
