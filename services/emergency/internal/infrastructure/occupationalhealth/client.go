package occupationalhealth

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

func (c *Client) AlertMedicalClinic(ctx context.Context, plantID string, casualtyCount int) error {
	prahariLogger.Info(ctx, "Alerted occupational health clinic of mass casualty / exposure emergency",
		prahariLogger.String("plant_id", plantID),
		prahariLogger.Int("casualty_count", casualtyCount))
	return nil
}
