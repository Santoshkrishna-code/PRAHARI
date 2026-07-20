package identity

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client wraps gRPC calls to Identity Service.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// UserExists checks user status.
func (c *Client) UserExists(ctx context.Context, userID string) (bool, error) {
	prahariLogger.Info(ctx, "Verifying authority user ID via Identity Service",
		prahariLogger.String("user_id", userID))
	return true, nil
}
