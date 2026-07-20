package training

import (
	"errors"
	"time"
)

// Training tracks safety courses certifications updates.
type Training struct {
	ID          string    `json:"id" db:"id"`
	WorkerID    string    `json:"worker_id" db:"worker_id"`
	CourseName  string    `json:"course_name" db:"course_name"`
	CompletedAt time.Time `json:"completed_at" db:"completed_at"`
	ExpiryDate  time.Time `json:"expiry_date" db:"expiry_date"`
}

// Validate checks domain invariants.
func (t *Training) Validate() error {
	if t.WorkerID == "" {
		return errors.New("worker ID is required")
	}
	if t.CourseName == "" {
		return errors.New("course name is required")
	}
	return nil
}
