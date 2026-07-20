package enrollment

import (
	"errors"
)

// Enrollment tracks course registration.
type Enrollment struct {
	ID         string `json:"id" db:"id"`
	TrainingID string `json:"training_id" db:"training_id"`
	TraineeID  string `json:"trainee_id" db:"trainee_id"`
	Status     string `json:"status" db:"status"` // ENROLLED, COMPLETED, FAILED
}

// Validate checks domain invariants.
func (e *Enrollment) Validate() error {
	if e.TrainingID == "" {
		return errors.New("training ID reference is required")
	}
	if e.TraineeID == "" {
		return errors.New("trainee ID reference is required")
	}
	return nil
}
