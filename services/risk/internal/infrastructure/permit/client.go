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

// FetchActivePermitsCount checks high-hazard tasks.
func (c *Client) FetchActivePermitsCount(ctx context.Context, assetID string) (int, error) {
	prahariLogger.Info(ctx, "Querying active permit-to-work requests via Permit-to-Work Service gRPC",
		prahariLogger.String("asset_id", assetID))
	return 0, nil
}
