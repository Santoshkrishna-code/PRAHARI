package technicalreview

import "time"

// Review represents engineering and technical verification for an MOC.
type Review struct {
	ID              string    `json:"id"`
	ChangeRequestID string    `json:"change_request_id"`
	Discipline      string    `json:"discipline"` // Mechanical, Electrical, Process, Instrumentation, Civil
	ReviewerID      string    `json:"reviewer_id"`
	Status          string    `json:"status"` // APPROVED, REJECTED, CHANGES_REQUESTED
	Findings        string    `json:"findings"`
	Conditions      string    `json:"conditions,omitempty"`
	ReviewedAt      time.Time `json:"reviewed_at"`
}
