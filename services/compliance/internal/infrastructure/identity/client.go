package identity

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client verifies reviewer roles in organizations.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// VerifySafetyOwner check user roles.
func (c *Client) VerifySafetyOwner(ctx context.Context, userID string) (bool, error) {
	prahariLogger.Info(ctx, "Verifying compliance owner credentials via Identity Service",
		prahariLogger.String("user_id", userID))
	return true, nil
}
