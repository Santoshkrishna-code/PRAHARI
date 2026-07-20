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

func (c *Client) LogRenewableEnergyRatio(ctx context.Context, plantID string, ratio float64) error {
	prahariLogger.Info(ctx, "Logging clean renewable energy mix ratio directly to ESG & Sustainability Service",
		prahariLogger.String("plant_id", plantID),
		prahariLogger.String("ratio", fmt.Sprintf("%.2f", ratio)))
	return nil
}
