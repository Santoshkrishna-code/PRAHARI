package communication

import "time"

// Message represents shift-wide notes, hazard alerts, or broadcast announcements.
type Message struct {
	ID          string    `json:"id"`
	ShiftID     string    `json:"shift_id"`
	SenderID    string    `json:"sender_id"`
	Channel     string    `json:"channel"` // EMAIL, SMS, DASHBOARD_BROADCAST, SIREN
	Content     string    `json:"content"`
	SentAt      time.Time `json:"sent_at"`
}
