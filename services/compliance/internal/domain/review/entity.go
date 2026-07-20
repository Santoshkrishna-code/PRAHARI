package review

import (
	"errors"
	"time"
)

// ComplianceReview logs verification checklists checks.
type ComplianceReview struct {
	ID             string    `json:"id" db:"id"`
	ComplianceID   string    `json:"compliance_id" db:"compliance_id"`
	ReviewerID     string    `json:"reviewer_id" db:"reviewer_id"`
	ReviewDate     time.Time `json:"review_date" db:"review_date"`
	NextReviewDate time.Time `json:"next_review_date" db:"next_review_date"`
	Notes          string    `json:"notes" db:"notes"`
}

// Validate checks domain invariants.
func (cr *ComplianceReview) Validate() error {
	if cr.ComplianceID == "" {
		return errors.New("compliance ID reference is required")
	}
	if cr.ReviewerID == "" {
		return errors.New("reviewer ID is required")
	}
	return nil
}
