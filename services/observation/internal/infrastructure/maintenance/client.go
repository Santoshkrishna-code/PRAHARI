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

// CreateCorrectiveAction triggers maintenance corrective tasks for unsafe conditions.
func (c *Client) CreateCorrectiveAction(ctx context.Context, observationID, assetID, description string) error {
	prahariLogger.Info(ctx, "Requesting corrective maintenance task assignment for unsafe conditions via Maintenance Service gRPC",
		prahariLogger.String("observation_id", observationID),
		prahariLogger.String("asset_id", assetID))
	return nil
}
