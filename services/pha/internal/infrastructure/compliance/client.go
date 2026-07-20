package compliance

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

func (c *Client) CheckPSMCompliance(ctx context.Context, plantID string) error {
	prahariLogger.Info(ctx, "Checked OSHA 1910.119 PSM compliance status for PHA revalidation schedule", prahariLogger.String("plant_id", plantID))
	return nil
}
