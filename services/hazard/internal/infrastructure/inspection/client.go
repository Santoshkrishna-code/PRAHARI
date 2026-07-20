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

// QueryFindings retrieves inspection findings checklists.
func (c *Client) QueryFindings(ctx context.Context, assetID string) ([]string, error) {
	prahariLogger.Info(ctx, "Querying safety audit findings via Inspection Service gRPC",
		prahariLogger.String("asset_id", assetID))
	return []string{}, nil
}
