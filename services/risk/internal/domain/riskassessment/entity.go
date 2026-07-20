package riskassessment

import (
	"errors"
	"time"
)

// RiskAssessment models ISO 31000 processes safety.
type RiskAssessment struct {
	ID             string    `json:"id" db:"id"`
	RiskID         string    `json:"risk_id" db:"risk_id"`
	AssessedByID   string    `json:"assessed_by_id" db:"assessed_by_id"`
	AssessmentDate time.Time `json:"assessment_date" db:"assessment_date"`
	Methodology    string    `json:"methodology" db:"methodology"` // HAZOP, Bow-Tie, JSA
	RiskScore      int       `json:"risk_score" db:"risk_score"`
}

// Validate checks domain invariants.
func (ra *RiskAssessment) Validate() error {
	if ra.RiskID == "" {
		return errors.New("risk register ID reference is required")
	}
	if ra.AssessedByID == "" {
		return errors.New("assessed by ID is required")
	}
	return nil
}
