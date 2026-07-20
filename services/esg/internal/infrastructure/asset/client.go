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

func (c *Client) FetchSolarGeneratorAssets(ctx context.Context, plantID string) ([]string, error) {
	prahariLogger.Info(ctx, "Querying solar generator clean energy assets from Asset Management Service",
		prahariLogger.String("plant_id", plantID))
	return []string{"ASSET-SOLAR-099"}, nil
}
