package consequence

import (
	"errors"
)

// Consequence maps values.
type Consequence struct {
	ID          string `json:"id" db:"id"`
	Value       int    `json:"value" db:"value"` // 1-5
	Name        string `json:"name" db:"name"`   // Insignificant, Minor, Moderate, Major, Catastrophic
	Description string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (c *Consequence) Validate() error {
	if c.Value < 1 || c.Value > 5 {
		return errors.New("consequence value must fall inside 1-5 range")
	}
	if c.Name == "" {
		return errors.New("consequence name is required")
	}
	return nil
}
