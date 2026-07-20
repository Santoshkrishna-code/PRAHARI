package trainingprogram

import (
	"errors"
)

// TrainingProgram defines structured tracks.
type TrainingProgram struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (tp *TrainingProgram) Validate() error {
	if tp.Name == "" {
		return errors.New("program name is required")
	}
	return nil
}
