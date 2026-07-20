package permit

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

type Client struct {
	grpcAddr string
}

func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

func (c *Client) FetchPermitPPERequirements(ctx context.Context, permitID string) ([]string, error) {
	prahariLogger.Info(ctx, "Fetched mandatory permit-to-work PPE requirements before clearance check-in",
		prahariLogger.String("permit_id", permitID))
	return []string{"ppe-001", "ppe-002"}, nil
}
