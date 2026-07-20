package department

import (
	"errors"
)

// Department details mapping.
type Department struct {
	ID   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

// Validate checks domain invariants.
func (d *Department) Validate() error {
	if d.ID == "" {
		return errors.New("department ID reference is required")
	}
	return nil
}
