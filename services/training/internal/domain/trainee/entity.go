package trainee

import (
	"errors"
)

// Trainee details.
type Trainee struct {
	ID         string `json:"id" db:"id"`
	TrainingID string `json:"training_id" db:"training_id"`
	UserID     string `json:"user_id" db:"user_id"`
}

// Validate checks domain invariants.
func (t *Trainee) Validate() error {
	if t.TrainingID == "" {
		return errors.New("training ID reference is required")
	}
	if t.UserID == "" {
		return errors.New("user ID reference is required")
	}
	return nil
}
