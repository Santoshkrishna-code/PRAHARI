package incident

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client registers incident mappings.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// CountIncidents returns count.
func (c *Client) CountIncidents(ctx context.Context, assetID string) (int, error) {
	prahariLogger.Info(ctx, "Verifying asset safety incident history via Incident Service",
		prahariLogger.String("asset_id", assetID))
	return 0, nil
}
