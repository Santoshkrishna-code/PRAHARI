package auditor

import (
	"errors"
)

// Auditor details.
type Auditor struct {
	ID      string `json:"id" db:"id"`
	AuditID string `json:"audit_id" db:"audit_id"`
	UserID  string `json:"user_id" db:"user_id"`
	Role    string `json:"role" db:"role"` // e.g. Lead Auditor, Co-Auditor
}

// Validate checks domain invariants.
func (a *Auditor) Validate() error {
	if a.AuditID == "" {
		return errors.New("audit ID reference is required")
	}
	if a.UserID == "" {
		return errors.New("user ID reference is required")
	}
	return nil
}
