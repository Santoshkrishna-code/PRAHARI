package effectiveness

import (
	"errors"
	"time"
)

// Effectiveness measures EHS safety learning impact scores.
type Effectiveness struct {
	ID              string    `json:"id" db:"id"`
	ObservationID   string    `json:"observation_id" db:"observation_id"`
	EvaluatorID     string    `json:"evaluator_id" db:"evaluator_id"`
	EvaluationDate  time.Time `json:"evaluation_date" db:"evaluation_date"`
	ImprovementRate int       `json:"improvement_rate" db:"improvement_rate"` // 1 to 5
	Notes           string    `json:"notes,omitempty" db:"notes"`
}

// Validate checks domain invariants.
func (e *Effectiveness) Validate() error {
	if e.ObservationID == "" {
		return errors.New("observation ID reference is required")
	}
	if e.EvaluatorID == "" {
		return errors.New("evaluator ID is required")
	}
	if e.ImprovementRate < 1 || e.ImprovementRate > 5 {
		return errors.New("improvement rate must fall inside a 1-5 range")
	}
	return nil
}
