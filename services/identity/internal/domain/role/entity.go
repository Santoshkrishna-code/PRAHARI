package role

import (
	"errors"
)

// Role defines access control profiles.
type Role struct {
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
}

// Validate checks parameters.
func (r *Role) Validate() error {
	if r.Name == "" {
		return errors.New("role name is required")
	}
	return nil
}
