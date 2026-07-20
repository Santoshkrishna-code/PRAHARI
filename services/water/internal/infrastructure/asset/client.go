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

func (c *Client) GetEquipmentSpecs(ctx context.Context, assetID string) (float64, error) {
	prahariLogger.Info(ctx, "Fetched treatment/pump equipment asset specs", prahariLogger.String("asset_id", assetID))
	return 100.0, nil
}
