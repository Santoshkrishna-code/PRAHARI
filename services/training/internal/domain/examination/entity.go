package examination

import (
	"errors"
)

// Examination maps written questions details.
type Examination struct {
	ID         string `json:"id" db:"id"`
	TrainingID string `json:"training_id" db:"training_id"`
	Topic      string `json:"topic" db:"topic"`
}

// Validate checks domain invariants.
func (e *Examination) Validate() error {
	if e.TrainingID == "" {
		return errors.New("training ID reference is required")
	}
	return nil
}
