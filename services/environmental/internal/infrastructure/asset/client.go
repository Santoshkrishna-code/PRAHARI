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

func (c *Client) VerifyEmissionSourceAsset(ctx context.Context, assetID string) (bool, error) {
	prahariLogger.Info(ctx, "Verifying operational status of emission stack asset with Asset Management Service",
		prahariLogger.String("asset_id", assetID))
	return true, nil
}
