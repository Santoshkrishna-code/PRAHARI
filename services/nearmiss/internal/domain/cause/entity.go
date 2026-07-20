package cause

import (
	"errors"
)

// Cause root causes classification taxonomy.
type Cause struct {
	ID          string `json:"id" db:"id"`
	NearMissID  string `json:"near_miss_id" db:"near_miss_id"`
	RootCause   string `json:"root_cause" db:"root_cause"` // Human factor, Equip failure, etc
	Description string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (c *Cause) Validate() error {
	if c.NearMissID == "" {
		return errors.New("near miss ID reference is required")
	}
	if c.RootCause == "" {
		return errors.New("root cause description is required")
	}
	return nil
}
