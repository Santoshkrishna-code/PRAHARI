package standard

import (
	"errors"
)

// Standard models ISO standard taxonomy items.
type Standard struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"` // ISO 45001, OSHA, Factory Act
	Description string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (s *Standard) Validate() error {
	if s.Name == "" {
		return errors.New("standard name is required")
	}
	return nil
}
