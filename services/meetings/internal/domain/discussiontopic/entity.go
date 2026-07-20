package discussiontopic

import "time"

// Topic represents a discussion topic raised during a meeting — safety alerts,
// lessons learned, hazard notifications, or operational updates.
type Topic struct {
	ID          string    `json:"id"`
	MeetingID   string    `json:"meeting_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	SourceType  string    `json:"source_type,omitempty"` // INCIDENT, HAZARD, NEAR_MISS, RISK, REGULATORY
	SourceRefID string    `json:"source_ref_id,omitempty"`
	RaisedBy    string    `json:"raised_by"`
	CreatedAt   time.Time `json:"created_at"`
}
