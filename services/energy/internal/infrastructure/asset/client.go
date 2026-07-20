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

func (c *Client) FetchEquipmentPowerRating(ctx context.Context, assetID string) (float64, error) {
	prahariLogger.Info(ctx, "Querying active baseline power ratings from Asset Management Service gRPC",
		prahariLogger.String("asset_id", assetID))
	return 45.0, nil // 45 KW active load
}
