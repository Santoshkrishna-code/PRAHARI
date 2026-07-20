package notification

import (
	"errors"
)

// Notification represents the central profile aggregate payload details.
type Notification struct {
	ID        string `json:"id"`
	Recipient string `json:"recipient"`
	Channel   string `json:"channel"` // e.g. "email", "sms", "slack"
	Content   string `json:"content"`
	Status    string `json:"status"` // e.g. "QUEUED", "SENT", "FAILED"
}

// Validate checks entity values parameters.
func (n *Notification) Validate() error {
	if n.ID == "" {
		return errors.New("notification ID is required")
	}
	if n.Recipient == "" {
		return errors.New("recipient coordinate value is required")
	}
	if n.Channel == "" {
		return errors.New("channel type is required")
	}
	return nil
}
