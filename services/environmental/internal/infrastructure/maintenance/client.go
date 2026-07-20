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

func (c *Client) RaiseAssetMaintenanceRequest(ctx context.Context, assetID string, desc string) error {
	prahariLogger.Info(ctx, "Raising asset preventive repair request with Maintenance Management Service due to emission parameters limits breach",
		prahariLogger.String("asset_id", assetID))
	return nil
}
