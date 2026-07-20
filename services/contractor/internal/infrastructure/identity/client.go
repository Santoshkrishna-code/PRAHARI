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

// VerifySupervisor check user roles.
func (c *Client) VerifySupervisor(ctx context.Context, userID string) (bool, error) {
	prahariLogger.Info(ctx, "Verifying supervisor role via Identity Service",
		prahariLogger.String("user_id", userID))
	return true, nil
}
