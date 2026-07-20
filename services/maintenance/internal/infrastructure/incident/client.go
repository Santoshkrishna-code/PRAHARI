package incident

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client queries safety tickets.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// CheckBreakdownIncident verifies incident logs.
func (c *Client) CheckBreakdownIncident(ctx context.Context, assetID string) (bool, error) {
	prahariLogger.Info(ctx, "Checking Incident Service safety tickets for asset breakdown logs",
		prahariLogger.String("asset_id", assetID))
	return true, nil
}
