package approval

import (
	"errors"
	"time"
)

// Role defines the level of authority required for a signature stage.
type Role string

const (
	RoleSupervisor     Role = "SUPERVISOR"
	RoleSafetyOfficer  Role = "SAFETY_OFFICER"
	RoleAreaAuthority  Role = "AREA_AUTHORITY"
	RolePlantManager   Role = "PLANT_MANAGER"
)

// Decision defines the outcome of an approval step.
type Decision string

const (
	DecisionPending  Decision = "PENDING"
	DecisionApproved Decision = "APPROVED"
	DecisionRejected Decision = "REJECTED"
)

// Approval represents an audit record of a user signing off at a lifecycle stage.
type Approval struct {
	ID            string     `json:"id" db:"id"`
	PermitID      string     `json:"permit_id" db:"permit_id"`
	ApproverID    string     `json:"approver_id" db:"approver_id"`
	ApproverRole  Role       `json:"approver_role" db:"approver_role"`
	Decision      Decision   `json:"decision" db:"decision"`
	Comments      string     `json:"comments,omitempty" db:"comments"`
	SignatureHash string     `json:"signature_hash,omitempty" db:"signature_hash"`
	DecidedAt     *time.Time `json:"decided_at,omitempty" db:"decided_at"`
	SequenceOrder int        `json:"sequence_order" db:"sequence_order"`
}

// Validate checks domain invariants for Approval.
func (a *Approval) Validate() error {
	if a.PermitID == "" {
		return errors.New("permit ID is required for approval")
	}
	if a.ApproverID == "" {
		return errors.New("approver ID is required")
	}
	if a.ApproverRole == "" {
		return errors.New("approver role is required")
	}
	return nil
}

// Approve applies signature fields to the approval step.
func (a *Approval) Approve(signature string) {
	now := time.Now()
	a.Decision = DecisionApproved
	a.SignatureHash = signature
	a.DecidedAt = &now
}

// Reject records rejection decision.
func (a *Approval) Reject(comments string) {
	now := time.Now()
	a.Decision = DecisionRejected
	a.Comments = comments
	a.DecidedAt = &now
}
