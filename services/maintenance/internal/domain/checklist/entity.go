package checklist

import (
	"errors"
)

// Checklist holds execution checkpoints verification tags list.
type Checklist struct {
	ID            string `json:"id" db:"id"`
	MaintenanceID string `json:"maintenance_id" db:"maintenance_id"`
	Name          string `json:"name" db:"name"`
	IsPassed      bool   `json:"is_passed" db:"is_passed"`
}

// Validate checks domain invariants.
func (c *Checklist) Validate() error {
	if c.MaintenanceID == "" {
		return errors.New("maintenance ID is required")
	}
	if c.Name == "" {
		return errors.New("checklist name is required")
	}
	return nil
}
