package notification

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) SendEnergyPeakSheddingAlert(ctx context.Context, engineerID string, message string) error {
	prahariLogger.Info(ctx, "Dispatching smart peak load shedding alert warnings via SMS/Email notifications",
		prahariLogger.String("engineer_id", engineerID))
	return nil
}
