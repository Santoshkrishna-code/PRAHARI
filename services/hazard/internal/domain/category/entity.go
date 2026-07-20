package category

import (
	"errors"
)

// Category classifications safety classifications.
type Category struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (c *Category) Validate() error {
	if c.Name == "" {
		return errors.New("category name is required")
	}
	return nil
}
