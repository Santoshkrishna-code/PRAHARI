package communication

import "time"

// Broadcast represents emergency situation updates sent to plant personnel or external authorities.
type Broadcast struct {
	ID          string    `json:"id"`
	EmergencyID string    `json:"emergency_id"`
	SenderID    string    `json:"sender_id"`
	Channel     string    `json:"channel"` // PUBLIC_ADDRESS, SMS, RADIO, EMAIL, SIREN
	Message     string    `json:"message"`
	SentAt      time.Time `json:"sent_at"`
}
