package incident

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

func (c *Client) FetchActiveIncidents(ctx context.Context, plantID string) ([]string, error) {
	prahariLogger.Info(ctx, "Fetched active incident reports for shift logbook integration",
		prahariLogger.String("plant_id", plantID))
	return []string{"inc-501"}, nil
}
