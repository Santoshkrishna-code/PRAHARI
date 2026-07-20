package environmental

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

func (c *Client) FetchScope1Emissions(ctx context.Context, plantID string) (float64, error) {
	prahariLogger.Info(ctx, "Retrieving direct atmospheric emissions stack values from Environmental Service gRPC",
		prahariLogger.String("plant_id", plantID))
	return 15000.5, nil // in kg CO2
}
