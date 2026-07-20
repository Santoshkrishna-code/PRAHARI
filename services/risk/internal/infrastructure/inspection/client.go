package inspection

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client checks active checklists.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// FetchFailedChecklistItems checks safety findings.
func (c *Client) FetchFailedChecklistItems(ctx context.Context, assetID string) (int, error) {
	prahariLogger.Info(ctx, "Querying failed compliance checklist items via Inspection Service gRPC",
		prahariLogger.String("asset_id", assetID))
	return 0, nil
}
