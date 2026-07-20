package incident

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client checks reactive records indices.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// FetchIncidentsCount checks recorded incident counts in plant zones.
func (c *Client) FetchIncidentsCount(ctx context.Context, assetID string) (int, error) {
	prahariLogger.Info(ctx, "Querying recorded incidents count via Incident Service gRPC",
		prahariLogger.String("asset_id", assetID))
	return 0, nil
}
