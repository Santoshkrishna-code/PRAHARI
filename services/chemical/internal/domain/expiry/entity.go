package expiry

import "time"

// Record tracks specific container expiry dates and alerts.
type Record struct {
	ID          string    `json:"id"`
	ContainerID string    `json:"container_id"`
	ExpiryDate  time.Time `json:"expiry_date"`
	AlertSent   bool      `json:"alert_sent"`
	AlertSentAt *time.Time `json:"alert_sent_at,omitempty"`
}
