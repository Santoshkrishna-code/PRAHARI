package communication

import "time"

// Record tracks stakeholder and plant notification for a change.
type Record struct {
	ID              string    `json:"id"`
	ChangeRequestID string    `json:"change_request_id"`
	RecipientGroup  string    `json:"recipient_group"`
	Subject         string    `json:"subject"`
	SentAt          time.Time `json:"sent_at"`
}
