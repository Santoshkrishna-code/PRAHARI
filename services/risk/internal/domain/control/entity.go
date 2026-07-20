package control

import (
	"errors"
)

// ControlType classifications.
type ControlType string

const (
	TypeElimination    ControlType = "ELIMINATION"
	TypeSubstitution   ControlType = "SUBSTITUTION"
	TypeEngineering    ControlType = "ENGINEERING"
	TypeAdministrative ControlType = "ADMINISTRATIVE"
	TypePPE            ControlType = "PPE"
)

// Control records mitigations barrier effectiveness ratings.
type Control struct {
	ID          string      `json:"id" db:"id"`
	RiskID      string      `json:"risk_id" db:"risk_id"`
	ControlType ControlType `json:"control_type" db:"control_type"`
	Description string      `json:"description" db:"description"`
	EffectValue int         `json:"effect_value" db:"effect_value"` // 1-5 effectiveness
}

// Validate checks domain invariants.
func (c *Control) Validate() error {
	if c.RiskID == "" {
		return errors.New("risk ID is required")
	}
	if c.Description == "" {
		return errors.New("control description cannot be empty")
	}
	if c.EffectValue < 1 || c.EffectValue > 5 {
		return errors.New("effectiveness rating must fall inside 1-5 range")
	}
	return nil
}
