package review

import (
	"errors"
	"time"
)

// AuditReview logs verification checklists checks.
type AuditReview struct {
	ID             string    `json:"id" db:"id"`
	AuditID        string    `json:"audit_id" db:"audit_id"`
	ReviewerID     string    `json:"reviewer_id" db:"reviewer_id"`
	ReviewDate     time.Time `json:"review_date" db:"review_date"`
	NextReviewDate time.Time `json:"next_review_date" db:"next_review_date"`
	Notes          string    `json:"notes" db:"notes"`
}

// Validate checks domain invariants.
func (ar *AuditReview) Validate() error {
	if ar.AuditID == "" {
		return errors.New("audit ID reference is required")
	}
	if ar.ReviewerID == "" {
		return errors.New("reviewer ID is required")
	}
	return nil
}
