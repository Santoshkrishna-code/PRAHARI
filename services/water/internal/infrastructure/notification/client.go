package notification

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) SendWaterAlertNotification(ctx context.Context, recipient, message string) error {
	prahariLogger.Info(ctx, "Sent Water management alert notification", prahariLogger.String("recipient", recipient))
	return nil
}
