package badge

import "time"

// Badge represents a physical RFID or temporary access card issued to a visitor.
type Badge struct {
	ID          string     `json:"id"`
	BadgeNumber string     `json:"badge_number"`
	RFIDCode    string     `json:"rfid_code,omitempty"`
	Status      string     `json:"status"` // AVAILABLE, ISSUED, LOST, EXPIRED
	IssuedTo    string     `json:"issued_to,omitempty"` // Visitor ID
	IssuedAt    *time.Time `json:"issued_at,omitempty"`
	ReturnedAt  *time.Time `json:"returned_at,omitempty"`
}
