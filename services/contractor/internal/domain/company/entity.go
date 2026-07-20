package company

import (
	"errors"
)

// Company holds general registry details for a contractor organization.
type Company struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	IsActive    bool   `json:"is_active" db:"is_active"`
}

// Validate checks domain invariants.
func (c *Company) Validate() error {
	if c.Name == "" {
		return errors.New("company name is required")
	}
	return nil
}
