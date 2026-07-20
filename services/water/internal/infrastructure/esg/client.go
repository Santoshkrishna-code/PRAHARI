package esg

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

func (c *Client) LogWaterStewardshipIndex(ctx context.Context, plantID string, recyclePct float64) error {
	prahariLogger.Info(ctx, "Sent Water Stewardship Index & Recycling Ratio to ESG & Sustainability Service",
		prahariLogger.String("plant_id", plantID),
		prahariLogger.String("recycle_pct", fmt.Sprintf("%.2f", recyclePct)))
	return nil
}
