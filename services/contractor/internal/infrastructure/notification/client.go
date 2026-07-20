package notification

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client formats Kafka notification payloads.
type Client struct{}

// NewClient instantiates Client.
func NewClient() *Client {
	return &Client{}
}

// BuildNotificationPayload formats event structures for Kafka notifications.
func (c *Client) BuildNotificationPayload(ctx context.Context, contractorID, eventType, recipientID, message string) map[string]string {
	prahariLogger.Info(ctx, "Formatting notification payload",
		prahariLogger.String("contractor_id", contractorID),
		prahariLogger.String("event_type", eventType))

	return map[string]string{
		"contractor_id": contractorID,
		"event_type":    eventType,
		"recipient_id":  recipientID,
		"message":       message,
	}
}
