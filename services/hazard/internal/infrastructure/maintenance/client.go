package maintenance

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client links contractor workorders allocations.
type Client struct {
	grpcAddr string
}

// NewClient instantiates Client.
func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

// CreateCorrectiveAction triggers maintenance corrective tasks.
func (c *Client) CreateCorrectiveAction(ctx context.Context, hazardID, assetID, description string) error {
	prahariLogger.Info(ctx, "Requesting corrective maintenance task assignment via Maintenance Service gRPC",
		prahariLogger.String("hazard_id", hazardID),
		prahariLogger.String("asset_id", assetID))
	return nil
}
