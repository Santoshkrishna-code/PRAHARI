package induction

import (
	"errors"
	"time"
)

// Induction tracks safety induction orientations completions.
type Induction struct {
	ID          string    `json:"id" db:"id"`
	WorkerID    string    `json:"worker_id" db:"worker_id"`
	LocationID  string    `json:"location_id" db:"location_id"`
	CompletedAt time.Time `json:"completed_at" db:"completed_at"`
	ValidUntil  time.Time `json:"valid_until" db:"valid_until"`
}

// Validate checks domain invariants.
func (i *Induction) Validate() error {
	if i.WorkerID == "" {
		return errors.New("worker ID reference is required")
	}
	if i.LocationID == "" {
		return errors.New("location ID reference is required")
	}
	return nil
}
