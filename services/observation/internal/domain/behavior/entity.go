package behavior

import (
	"errors"
)

// Behavior maps safe or unsafe behaviors observed.
type Behavior struct {
	ID          string `json:"id" db:"id"`
	ObservationID string `json:"observation_id" db:"observation_id"`
	Category    string `json:"category" db:"category"` // e.g. PPE Compliance, Permit Compliance
	IsSafe      bool   `json:"is_safe" db:"is_safe"`
	Description string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (b *Behavior) Validate() error {
	if b.ObservationID == "" {
		return errors.New("observation ID is required")
	}
	if b.Category == "" {
		return errors.New("behavior category is required")
	}
	return nil
}
