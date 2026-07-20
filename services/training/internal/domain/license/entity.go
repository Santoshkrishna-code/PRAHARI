package license

import (
	"errors"
	"time"
)

// License details statutory permissions (heavy vehicles, high voltage).
type License struct {
	ID         string    `json:"id" db:"id"`
	TraineeID  string    `json:"trainee_id" db:"trainee_id"`
	Issuer     string    `json:"issuer" db:"issuer"`
	ValidUntil time.Time `json:"valid_until" db:"valid_until"`
}

// Validate checks domain invariants.
func (l *License) Validate() error {
	if l.TraineeID == "" {
		return errors.New("trainee ID is required")
	}
	return nil
}
