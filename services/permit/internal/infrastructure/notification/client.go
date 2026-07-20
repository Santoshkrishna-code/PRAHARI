package notification

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client constructs Kafka notification payloads.
type Client struct{}

// NewClient instantiates Client.
func NewClient() *Client {
	return &Client{}
}

// BuildNotificationPayload formats event structures for Kafka notifications.
func (c *Client) BuildNotificationPayload(ctx context.Context, permitID, eventType, recipientID, message string) map[string]string {
	prahariLogger.Info(ctx, "Formatting notification payload",
		prahariLogger.String("permit_id", permitID),
		prahariLogger.String("event_type", eventType))

	return map[string]string{
		"permit_id":    permitID,
		"event_type":   eventType,
		"recipient_id": recipientID,
		"message":      message,
	}
}
