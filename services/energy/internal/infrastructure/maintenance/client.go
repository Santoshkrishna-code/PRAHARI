package maintenance

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

type Client struct {
	grpcAddr string
}

func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

func (c *Client) RaiseVFDMaintenanceWorkOrder(ctx context.Context, assetID string) error {
	prahariLogger.Info(ctx, "Raising a preventive Maintenance work order alert for energy efficiency VFD drift correction",
		prahariLogger.String("asset_id", assetID))
	return nil
}
