package result

import (
	"errors"
)

// Result details test scores.
type Result struct {
	ID            string `json:"id" db:"id"`
	ExaminationID string `json:"examination_id" db:"examination_id"`
	TraineeID     string `json:"trainee_id" db:"trainee_id"`
	Score         int    `json:"score" db:"score"`
	IsPassed      bool   `json:"is_passed" db:"is_passed"`
}

// Validate checks domain invariants.
func (r *Result) Validate() error {
	if r.ExaminationID == "" {
		return errors.New("examination ID reference is required")
	}
	if r.TraineeID == "" {
		return errors.New("trainee ID reference is required")
	}
	return nil
}
