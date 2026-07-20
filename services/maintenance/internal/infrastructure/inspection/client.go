package inspection

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client checks active inspection checklists.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// QueryFindings returns findings list.
func (c *Client) QueryFindings(ctx context.Context, assetID string) ([]string, error) {
	prahariLogger.Info(ctx, "Retrieving inspection recommendations from Inspection Service gRPC",
		prahariLogger.String("asset_id", assetID))
	return []string{}, nil
}
