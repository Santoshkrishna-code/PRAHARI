package acknowledgement

import "time"

// Acknowledgement represents a digital signature confirming that an attendee has
// understood the safety topics discussed during a meeting or toolbox talk.
type Acknowledgement struct {
	ID           string    `json:"id"`
	MeetingID    string    `json:"meeting_id"`
	AttendeeID   string    `json:"attendee_id"`
	SignatureURL string    `json:"signature_url,omitempty"`
	AcknowledgedAt time.Time `json:"acknowledged_at"`
	IPAddress    string    `json:"ip_address,omitempty"`
	DeviceID     string    `json:"device_id,omitempty"`
}
