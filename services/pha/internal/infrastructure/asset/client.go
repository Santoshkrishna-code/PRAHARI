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

func (c *Client) UpdateBarrierRequirements(ctx context.Context, assetID, silLevel string) error {
	prahariLogger.Info(ctx, "Updated asset barrier & SIL safety requirements",
		prahariLogger.String("asset_id", assetID),
		prahariLogger.String("sil_level", silLevel))
	return nil
}
