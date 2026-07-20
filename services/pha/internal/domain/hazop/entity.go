package hazop

import "time"

// Session represents a HAZOP review session record.
type Session struct {
	ID          string    `json:"id"`
	StudyID     string    `json:"study_id"`
	SessionDate time.Time `json:"session_date"`
	DurationHrs float64   `json:"duration_hrs"`
	Attendees   string    `json:"attendees"` // Comma-separated list of user IDs
	Notes       string    `json:"notes"`
	CreatedAt   time.Time `json:"created_at"`
}
