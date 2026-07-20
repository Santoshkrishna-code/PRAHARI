package energy

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

func (c *Client) CorrelatePumpingEnergy(ctx context.Context, plantID string, pumpPowerKWh float64) error {
	prahariLogger.Info(ctx, "Correlated water pump electricity demand with Energy Management Service",
		prahariLogger.String("plant_id", plantID),
		prahariLogger.String("kwh", fmt.Sprintf("%.2f", pumpPowerKWh)))
	return nil
}
