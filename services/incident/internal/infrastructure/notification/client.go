package notification

import (
	"context"

	prahariLogger "prahari/shared/logger"
)

// Client provides helpers for constructing well-formed Kafka event payloads
// that the Notification Service consumes. The Incident Service never calls
// the Notification Service directly — all communication is event-driven.
type Client struct{}

// NewClient constructs a Notification client.
func NewClient() *Client {
	return &Client{}
}

// BuildNotificationPayload constructs a Kafka event payload for the Notification Service.
func (c *Client) BuildNotificationPayload(ctx context.Context, incidentID, eventType, recipientID, message string) map[string]string {
	prahariLogger.Info(ctx, "Building notification payload for Kafka event",
		prahariLogger.String("incident_id", incidentID),
		prahariLogger.String("event_type", eventType))

	return map[string]string{
		"incident_id":  incidentID,
		"event_type":   eventType,
		"recipient_id": recipientID,
		"message":      message,
	}
}
