package maintenance

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

func (c *Client) FetchActiveWorkOrders(ctx context.Context, plantID string) ([]string, error) {
	prahariLogger.Info(ctx, "Fetched active maintenance work orders for shift log integration",
		prahariLogger.String("plant_id", plantID))
	return []string{"wo-3001", "wo-3002"}, nil
}
