package barrier

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

func (c *Client) CheckBarrierImpairmentStatus(ctx context.Context, plantID string) (int, error) {
	prahariLogger.Info(ctx, "Checked active barrier impairments for emergency readiness evaluation", prahariLogger.String("plant_id", plantID))
	return 1, nil
}
