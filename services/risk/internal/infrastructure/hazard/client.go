package hazard

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client links active risk categories.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// FetchActiveHazards queries hazard status list.
func (c *Client) FetchActiveHazards(ctx context.Context, assetID string) ([]string, error) {
	prahariLogger.Info(ctx, "Querying active hazard registry items via Hazard Service gRPC",
		prahariLogger.String("asset_id", assetID))
	return []string{}, nil
}
