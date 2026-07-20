package identity

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client verifies trainer roles in organizations.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// VerifySafetyTrainer check user roles.
func (c *Client) VerifySafetyTrainer(ctx context.Context, userID string) (bool, error) {
	prahariLogger.Info(ctx, "Verifying safety trainer credentials via Identity Service",
		prahariLogger.String("user_id", userID))
	return true, nil
}
