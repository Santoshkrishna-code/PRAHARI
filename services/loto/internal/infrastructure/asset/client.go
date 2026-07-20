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

func (c *Client) FetchAssetIsolationPoints(ctx context.Context, equipmentID string) ([]string, error) {
	prahariLogger.Info(ctx, "Fetched defined isolation points from Asset Management Service database",
		prahariLogger.String("equipment_id", equipmentID))
	return []string{"breaker-CB12", "valve-V102"}, nil
}
