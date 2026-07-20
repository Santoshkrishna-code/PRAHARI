package assessment

import (
	"errors"
)

// Assessment evaluates field competency checks.
type Assessment struct {
	ID          string `json:"id" db:"id"`
	TrainingID  string `json:"training_id" db:"training_id"`
	TraineeID   string `json:"trainee_id" db:"trainee_id"`
	Score       int    `json:"score" db:"score"`
	IsPassed    bool   `json:"is_passed" db:"is_passed"`
}

// Validate checks domain invariants.
func (a *Assessment) Validate() error {
	if a.TrainingID == "" {
		return errors.New("training ID reference is required")
	}
	if a.TraineeID == "" {
		return errors.New("trainee ID reference is required")
	}
	return nil
}
