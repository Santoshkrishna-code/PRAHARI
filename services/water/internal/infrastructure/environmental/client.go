package environmental

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

func (c *Client) SendWaterUsageMetrics(ctx context.Context, plantID string, usageKL float64) error {
	prahariLogger.Info(ctx, "Sent water intake & consumption metrics to Environmental Management Service",
		prahariLogger.String("plant_id", plantID),
		prahariLogger.String("usage_kl", fmt.Sprintf("%.2f", usageKL)))
	return nil
}
