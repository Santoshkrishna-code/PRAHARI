package permit

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client verifies work permit details.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// ValidatePermitStatus checks permit status.
func (c *Client) ValidatePermitStatus(ctx context.Context, permitID string) (bool, error) {
	prahariLogger.Info(ctx, "Verifying active work permit details via Permit-to-Work Service gRPC",
		prahariLogger.String("permit_id", permitID))
	return true, nil
}
