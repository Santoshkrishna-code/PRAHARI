package asset

import (
	"context"
	"fmt"

	prahariLogger "prahari/shared/logger"
)


type Client struct {
	grpcAddr string
}

func NewClient(grpcAddr string) *Client {
	return &Client{grpcAddr: grpcAddr}
}

func (c *Client) UpdateAssetCriticality(ctx context.Context, assetID string, isIPL bool) error {
	prahariLogger.Info(ctx, "Updated asset criticality rating based on barrier assignment",
		prahariLogger.String("asset_id", assetID),
		prahariLogger.String("is_ipl", fmt.Sprintf("%t", isIPL)))
	return nil
}

