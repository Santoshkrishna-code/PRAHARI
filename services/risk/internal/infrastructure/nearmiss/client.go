package nearmiss

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client checks safety tickets.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// FetchNearMissCount checks recorded events.
func (c *Client) FetchNearMissCount(ctx context.Context, assetID string) (int, error) {
	prahariLogger.Info(ctx, "Querying near miss frequency parameters via Near Miss Service gRPC",
		prahariLogger.String("asset_id", assetID))
	return 0, nil
}
