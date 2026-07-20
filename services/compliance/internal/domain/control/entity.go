package control

import (
	"errors"
)

// Control maps compliance check procedures.
type Control struct {
	ID           string `json:"id" db:"id"`
	ComplianceID string `json:"compliance_id" db:"compliance_id"`
	Description  string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (c *Control) Validate() error {
	if c.ComplianceID == "" {
		return errors.New("compliance ID is required")
	}
	if c.Description == "" {
		return errors.New("control description cannot be empty")
	}
	return nil
}
