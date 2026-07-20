package location

import (
	"errors"
)

// Location models hierarchical plant nodes layout structures.
type Location struct {
	ID          string `json:"id" db:"id"`
	ParentID    string `json:"parent_id,omitempty" db:"parent_id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	IsActive    bool   `json:"is_active" db:"is_active"`
}

// Validate checks domain invariants.
func (l *Location) Validate() error {
	if l.Name == "" {
		return errors.New("location name is required")
	}
	return nil
}
