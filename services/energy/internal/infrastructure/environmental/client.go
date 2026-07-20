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

func (c *Client) SendEnergyEmissions(ctx context.Context, plantID string, co2Kg float64) error {
	prahariLogger.Info(ctx, "Logging indirect Scope 2 carbon footprint metrics directly to Environmental Service gRPC",
		prahariLogger.String("plant_id", plantID),
		prahariLogger.String("co2_kg", fmt.Sprintf("%.2f", co2Kg)))
	return nil
}
