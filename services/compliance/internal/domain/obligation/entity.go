package obligation

import (
	"errors"
	"time"
)

// Obligation models statutory compliance checklists items.
type Obligation struct {
	ID             string    `json:"id" db:"id"`
	ComplianceID   string    `json:"compliance_id" db:"compliance_id"`
	RegulationID   string    `json:"regulation_id" db:"regulation_id"`
	StandardID     string    `json:"standard_id" db:"standard_id"`
	DueDate        time.Time `json:"due_date" db:"due_date"`
	ExpirationDate time.Time `json:"expiration_date" db:"expiration_date"`
}

// Validate checks domain invariants.
func (o *Obligation) Validate() error {
	if o.ComplianceID == "" {
		return errors.New("compliance ID reference is required")
	}
	if o.RegulationID == "" {
		return errors.New("regulation ID reference is required")
	}
	return nil
}
