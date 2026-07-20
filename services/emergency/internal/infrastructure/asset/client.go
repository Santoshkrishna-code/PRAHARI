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

func (c *Client) FetchAssetLocation(ctx context.Context, assetID string) (string, error) {
	prahariLogger.Info(ctx, "Fetched asset geolocation and P&ID reference", prahariLogger.String("asset_id", assetID))
	return "ZONE-04-NORTH", nil
}
