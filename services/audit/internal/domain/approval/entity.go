package approval

import (
	"errors"
	"time"
)

// Approval registers signatures on audit milestones.
type Approval struct {
	ID           string    `json:"id" db:"id"`
	AuditID      string    `json:"audit_id" db:"audit_id"`
	ApproverID   string    `json:"approver_id" db:"approver_id"`
	ApprovedDate time.Time `json:"approved_date" db:"approved_date"`
	Signature    string    `json:"signature" db:"signature"` // Digital verification signature
}

// Validate checks domain invariants.
func (a *Approval) Validate() error {
	if a.AuditID == "" {
		return errors.New("audit ID is required")
	}
	if a.Signature == "" {
		return errors.New("digital signature string is required")
	}
	return nil
}
