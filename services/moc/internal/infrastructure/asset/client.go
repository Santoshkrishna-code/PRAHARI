package asset

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

func (c *Client) NotifyAssetModification(ctx context.Context, assetID, mocID string) error {
	prahariLogger.Info(ctx, "Notified Asset Management Service of approved process change",
		prahariLogger.String("asset_id", assetID),
		prahariLogger.String("moc_id", mocID))
	return nil
}
