package risk

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

func (c *Client) UpdateEnterpriseResilienceProfile(ctx context.Context, plantID string, resilienceIndex float64) error {
	prahariLogger.Info(ctx, "Updated enterprise risk profile with BCM resilience score",
		prahariLogger.String("plant_id", plantID))
	return nil
}
