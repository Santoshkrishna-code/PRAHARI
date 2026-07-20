package approval

import (
	"errors"
	"time"
)

// Approval registers signatures on assessment milestones.
type Approval struct {
	ID           string    `json:"id" db:"id"`
	RiskID       string    `json:"risk_id" db:"risk_id"`
	ApproverID   string    `json:"approver_id" db:"approver_id"`
	ApprovedDate time.Time `json:"approved_date" db:"approved_date"`
	Signature    string    `json:"signature" db:"signature"` // Digital verification signature
}

// Validate checks domain invariants.
func (a *Approval) Validate() error {
	if a.RiskID == "" {
		return errors.New("risk ID is required")
	}
	if a.Signature == "" {
		return errors.New("digital signature string is required")
	}
	return nil
}
