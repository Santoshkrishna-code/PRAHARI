package auditee

import (
	"errors"
)

// Auditee details.
type Auditee struct {
	ID      string `json:"id" db:"id"`
	AuditID string `json:"audit_id" db:"audit_id"`
	UserID  string `json:"user_id" db:"user_id"`
}

// Validate checks domain invariants.
func (a *Auditee) Validate() error {
	if a.AuditID == "" {
		return errors.New("audit ID reference is required")
	}
	if a.UserID == "" {
		return errors.New("user ID reference is required")
	}
	return nil
}
