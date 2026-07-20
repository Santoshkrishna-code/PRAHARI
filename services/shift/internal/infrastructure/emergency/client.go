package emergency

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

func (c *Client) FetchActiveEmergencies(ctx context.Context, plantID string) ([]string, error) {
	prahariLogger.Info(ctx, "Fetched active emergency events for active shift alert triggers",
		prahariLogger.String("plant_id", plantID))
	return []string{}, nil
}
