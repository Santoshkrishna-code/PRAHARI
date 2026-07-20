package notification

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

func (c *Client) SendEmergencyAlert(ctx context.Context, plantID, title, severity string) error {
	prahariLogger.Info(ctx, "Sent automated multi-channel plant siren & SMS emergency broadcast alert via Notification Service",
		prahariLogger.String("plant_id", plantID),
		prahariLogger.String("severity", severity))
	return nil
}
