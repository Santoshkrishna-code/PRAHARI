package attendance

import (
	"errors"
	"time"
)

// Attendance logs verification checkers.
type Attendance struct {
	ID           string    `json:"id" db:"id"`
	TrainingID   string    `json:"training_id" db:"training_id"`
	TraineeID    string    `json:"trainee_id" db:"trainee_id"`
	AttendedDate time.Time `json:"attended_date" db:"attended_date"`
	IsPresent    bool      `json:"is_present" db:"is_present"`
}

// Validate checks domain invariants.
func (a *Attendance) Validate() error {
	if a.TrainingID == "" {
		return errors.New("training ID reference is required")
	}
	if a.TraineeID == "" {
		return errors.New("trainee ID reference is required")
	}
	return nil
}
