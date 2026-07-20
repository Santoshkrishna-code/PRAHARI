package medical

import (
	"errors"
	"time"
)

// Medical clearance tracking.
type Medical struct {
	ID          string    `json:"id" db:"id"`
	WorkerID    string    `json:"worker_id" db:"worker_id"`
	EvaluatedAt time.Time `json:"evaluated_at" db:"evaluated_at"`
	ExpiryDate  time.Time `json:"expiry_date" db:"expiry_date"`
	IsFit       bool      `json:"is_fit" db:"is_fit"`
}

// Validate checks domain invariants.
func (m *Medical) Validate() error {
	if m.WorkerID == "" {
		return errors.New("worker ID reference is required")
	}
	return nil
}
