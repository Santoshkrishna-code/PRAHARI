package finding

import (
	"errors"
)

// Finding classifies gaps.
type Finding struct {
	ID          string `json:"id" db:"id"`
	AuditID     string `json:"audit_id" db:"audit_id"`
	FindingType string `json:"finding_type" db:"finding_type"` // Observation, Non-Conformity NCR
	Description string `json:"description" db:"description"`
}

// Validate checks domain invariants.
func (f *Finding) Validate() error {
	if f.AuditID == "" {
		return errors.New("audit ID reference is required")
	}
	if f.Description == "" {
		return errors.New("finding description cannot be empty")
	}
	return nil
}
