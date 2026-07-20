package likelihood

import (
	"errors"
)

// Likelihood maps values.
type Likelihood struct {
	ID          string `json:"id" db:"id"`
	Value       int    `json:"value" db:"value"` // 1-5
	Name        string `json:"name" db:"name"`   // Rare, Unlikely, Possible, Likely, Almost Certain
	Description string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (l *Likelihood) Validate() error {
	if l.Value < 1 || l.Value > 5 {
		return errors.New("likelihood value must fall inside 1-5 range")
	}
	if l.Name == "" {
		return errors.New("likelihood name is required")
	}
	return nil
}
