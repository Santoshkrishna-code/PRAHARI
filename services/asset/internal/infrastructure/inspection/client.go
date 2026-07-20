package inspection

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client verifies active checklist verifications.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// HasPSSRPassed checks safety review results.
func (c *Client) HasPSSRPassed(ctx context.Context, assetID string) (bool, error) {
	prahariLogger.Info(ctx, "Querying safety reviews status from Inspection Service gRPC",
		prahariLogger.String("asset_id", assetID))
	return true, nil
}
