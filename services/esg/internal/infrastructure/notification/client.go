package notification

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

type Client struct{}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) SendESGTargetAlert(ctx context.Context, officerID string, message string) error {
	prahariLogger.Info(ctx, "Dispatching corporate sustainability carbon milestone alert warning via SMS/Email",
		prahariLogger.String("officer_id", officerID))
	return nil
}
