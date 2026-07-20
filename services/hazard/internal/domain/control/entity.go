package control

import (
	"errors"
)

// ControlType hierarchy classification (Elimination, Substitution, Engineering, Administration, PPE).
type ControlType string

const (
	TypeElimination    ControlType = "ELIMINATION"
	TypeSubstitution   ControlType = "SUBSTITUTION"
	TypeEngineering    ControlType = "ENGINEERING"
	TypeAdministrative ControlType = "ADMINISTRATIVE"
	TypePPE            ControlType = "PPE"
)

// Control details safety controls applied.
type Control struct {
	ID          string      `json:"id" db:"id"`
	HazardID    string      `json:"hazard_id" db:"hazard_id"`
	ControlType ControlType `json:"control_type" db:"control_type"`
	Description string      `json:"description" db:"description"`
	IsActive    bool        `json:"is_active" db:"is_active"`
}

// Validate checks domain invariants.
func (c *Control) Validate() error {
	if c.HazardID == "" {
		return errors.New("hazard ID reference is required")
	}
	if c.Description == "" {
		return errors.New("control description cannot be empty")
	}
	return nil
}
