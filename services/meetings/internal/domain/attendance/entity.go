package attendance

import "time"

// Record represents an individual attendee's presence record for a meeting.
type Record struct {
	ID          string     `json:"id"`
	MeetingID   string     `json:"meeting_id"`
	AttendeeID  string     `json:"attendee_id"`
	AttendeeName string   `json:"attendee_name"`
	CheckInAt   time.Time  `json:"check_in_at"`
	CheckOutAt  *time.Time `json:"check_out_at,omitempty"`
	Method      string     `json:"method"` // MANUAL, QR_CODE, BIOMETRIC, DIGITAL_SIGN
	Verified    bool       `json:"verified"`
}
