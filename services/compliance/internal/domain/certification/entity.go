package certification

import (
	"errors"
	"time"
)

// Certification tracks active credentials/licenses validity parameters.
type Certification struct {
	ID          string    `json:"id" db:"id"`
	Issuer      string    `json:"issuer" db:"issuer"`
	ValidUntil  time.Time `json:"valid_until" db:"valid_until"`
	Status      string    `json:"status" db:"status"` // ACTIVE, EXPIRED
}

// Validate checks domain invariants.
func (c *Certification) Validate() error {
	if c.Issuer == "" {
		return errors.New("certification issuer is required")
	}
	return nil
}
