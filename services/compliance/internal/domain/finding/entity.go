package finding

import (
	"errors"
)

// Finding records non-compliant EHS violations audit checklist gaps.
type Finding struct {
	ID           string `json:"id" db:"id"`
	ComplianceID string `json:"compliance_id" db:"compliance_id"`
	Severity     string `json:"severity" db:"severity"` // Minor, Major
	Description  string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (f *Finding) Validate() error {
	if f.ComplianceID == "" {
		return errors.New("compliance ID reference is required")
	}
	if f.Description == "" {
		return errors.New("finding description cannot be empty")
	}
	return nil
}
