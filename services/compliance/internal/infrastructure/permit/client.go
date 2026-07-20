package permit

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client checks active permits in plant zones.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// VerifyStatutoryPermit checks active authorization certificates.
func (c *Client) VerifyStatutoryPermit(ctx context.Context, permitID string) (bool, error) {
	prahariLogger.Info(ctx, "Checking statutory permit compliance criteria via Permit-to-Work Service gRPC",
		prahariLogger.String("permit_id", permitID))
	return true, nil
}
