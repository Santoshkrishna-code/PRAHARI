package document

import "time"

// Document represents a controlled enterprise document aggregate root.
type Document struct {
	ID             string     `json:"id"`
	DocumentNumber string     `json:"document_number"`
	PlantID        string     `json:"plant_id"`
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	CategoryID     string     `json:"category_id"`
	DocumentType   string     `json:"document_type"` // SOP, WORK_INSTRUCTION, POLICY, STANDARD, PROCEDURE, FORM, PID, MSDS
	CurrentVersion string     `json:"current_version"`
	Status         string     `json:"status"` // Draft, Review, Approval, Published, Controlled Distribution, Periodic Review, Revision, Superseded, Archived, Rejected, Withdrawn
	CheckedOutBy   string     `json:"checked_out_by,omitempty"`
	CheckedOutAt   *time.Time `json:"checked_out_at,omitempty"`
	OwnerID        string     `json:"owner_id"`
	ReviewCycleM   int        `json:"review_cycle_months"`
	NextReviewAt   *time.Time `json:"next_review_at,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}
