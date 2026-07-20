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

func (c *Client) FetchEnergyEfficiencyMaintenanceUpgrades(ctx context.Context, plantID string) ([]string, error) {
	prahariLogger.Info(ctx, "Retrieving energy conservation work orders history details from Maintenance Service gRPC",
		prahariLogger.String("plant_id", plantID))
	return []string{}, nil
}
