package notification

import "time"

// Alert represents automated emergency escalation notifications.
type Alert struct {
	ID          string    `json:"id"`
	EmergencyID string    `json:"emergency_id"`
	Recipient   string    `json:"recipient"`
	AlertType   string    `json:"alert_type"` // EVACUATION, ALL_CLEAR, COMMAND_UPDATE
	Status      string    `json:"status"`
	SentAt      time.Time `json:"sent_at"`
}
