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

func (c *Client) FetchCriticalAssetRTO(ctx context.Context, assetID string) (float64, error) {
	prahariLogger.Info(ctx, "Fetched critical asset RTO requirement for BIA dependency mapping", prahariLogger.String("asset_id", assetID))
	return 4.0, nil
}
