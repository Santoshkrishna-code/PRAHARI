package role

import (
	"errors"
)

// Role defines plant duties.
type Role struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (r *Role) Validate() error {
	if r.Name == "" {
		return errors.New("role name is required")
	}
	return nil
}
