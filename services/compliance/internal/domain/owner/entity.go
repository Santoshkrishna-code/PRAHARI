package owner

import (
	"errors"
)

// Owner assigns compliance responsibility.
type Owner struct {
	ID           string `json:"id" db:"id"`
	ComplianceID string `json:"compliance_id" db:"compliance_id"`
	UserID       string `json:"user_id" db:"user_id"`
}

// Validate checks domain invariants.
func (o *Owner) Validate() error {
	if o.ComplianceID == "" {
		return errors.New("compliance ID reference is required")
	}
	if o.UserID == "" {
		return errors.New("user ID reference is required")
	}
	return nil
}
