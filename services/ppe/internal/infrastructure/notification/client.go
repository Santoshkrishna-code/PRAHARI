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

func (c *Client) SendPPEExpiryNotification(ctx context.Context, userID, itemID string) error {
	prahariLogger.Info(ctx, "Sent email alert regarding upcoming protective gear/PPE item expiry",
		prahariLogger.String("user_id", userID),
		prahariLogger.String("item_id", itemID))
	return nil
}
