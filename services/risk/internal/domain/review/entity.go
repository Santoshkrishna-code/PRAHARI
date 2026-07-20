package review

import (
	"errors"
	"time"
)

// RiskReview monitors risk reassessment schedules.
type RiskReview struct {
	ID             string    `json:"id" db:"id"`
	RiskID         string    `json:"risk_id" db:"risk_id"`
	ReviewerID     string    `json:"reviewer_id" db:"reviewer_id"`
	ReviewDate     time.Time `json:"review_date" db:"review_date"`
	NextReviewDate time.Time `json:"next_review_date" db:"next_review_date"`
	Notes          string    `json:"notes" db:"notes"`
}

// Validate checks domain invariants.
func (r *RiskReview) Validate() error {
	if r.RiskID == "" {
		return errors.New("risk ID reference is required")
	}
	if r.ReviewerID == "" {
		return errors.New("reviewer ID is required")
	}
	return nil
}
