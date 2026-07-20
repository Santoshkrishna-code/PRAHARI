package fitnessassessment

import (
	"errors"
	"time"
)

// FitnessAssessment records evaluations of a worker's fitness for duty.
type FitnessAssessment struct {
	ID              string    `json:"id" db:"id"`
	HealthProfileID string    `json:"health_profile_id" db:"health_profile_id"`
	AssessmentDate  time.Time `json:"assessment_date" db:"assessment_date"`
	EvaluatorID     string    `json:"evaluator_id" db:"evaluator_id"` // Physician ID
	ResultCode      string    `json:"result_code" db:"result_code"`   // "FIT", "FIT_WITH_RESTRICTIONS", "UNFIT_TEMPORARY", "UNFIT_PERMANENT"
	Notes           string    `json:"notes" db:"notes"`
	NextReviewDate  time.Time `json:"next_review_date" db:"next_review_date"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks assessment details.
func (f *FitnessAssessment) Validate() error {
	if f.HealthProfileID == "" {
		return errors.New("health profile reference is required")
	}
	if f.ResultCode == "" {
		return errors.New("fitness result code is required")
	}
	if f.EvaluatorID == "" {
		return errors.New("evaluator ID is required")
	}
	return nil
}
