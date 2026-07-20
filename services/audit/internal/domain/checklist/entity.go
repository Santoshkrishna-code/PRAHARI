package checklist

import (
	"errors"
)

// Checklist defines items scope.
type Checklist struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (c *Checklist) Validate() error {
	if c.Name == "" {
		return errors.New("checklist name is required")
	}
	return nil
}
