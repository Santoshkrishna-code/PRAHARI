package documentreview

import "time"

// Review tracks periodic SME / Technical reviews of controlled documents.
type Review struct {
	ID         string    `json:"id"`
	DocumentID string    `json:"document_id"`
	ReviewerID string    `json:"reviewer_id"`
	Status     string    `json:"status"` // PENDING, APPROVED, REJECTED
	Comments   string    `json:"comments"`
	ReviewedAt time.Time `json:"reviewed_at"`
}
