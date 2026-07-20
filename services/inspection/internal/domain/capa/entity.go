package capa

import (
	"errors"
	"time"
)

// ActionType defines corrective or preventive category classifications.
type ActionType string

const (
	TypeCorrective ActionType = "CORRECTIVE"
	TypePreventive ActionType = "PREVENTIVE"
)

// Status defines tracking stages.
type Status string

const (
	StatusOpen      Status = "OPEN"
	StatusCompleted Status = "COMPLETED"
	StatusVerified  Status = "VERIFIED"
)

// CAPA represents a corrective and preventive action generated from critical inspection findings.
type CAPA struct {
	ID          string     `json:"id" db:"id"`
	InspectionID string    `json:"inspection_id" db:"inspection_id"`
	FindingID   string     `json:"finding_id" db:"finding_id"`
	ActionType  ActionType `json:"action_type" db:"action_type"`
	Description string     `json:"description" db:"description"`
	AssigneeID  string     `json:"assignee_id" db:"assignee_id"`
	DueDate     time.Time  `json:"due_date" db:"due_date"`
	CompletedAt *time.Time `json:"completed_at,omitempty" db:"completed_at"`
	Status      Status     `json:"status" db:"status"`
	VerifiedBy  string     `json:"verified_by,omitempty" db:"verified_by"`
	VerifiedAt  *time.Time `json:"verified_at,omitempty" db:"verified_at"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
}

// Validate checks domain invariants.
func (c *CAPA) Validate() error {
	if c.InspectionID == "" {
		return errors.New("inspection ID is required for CAPA")
	}
	if c.Description == "" {
		return errors.New("CAPA description is required")
	}
	if c.AssigneeID == "" {
		return errors.New("assignee user ID is required")
	}
	if c.DueDate.IsZero() {
		return errors.New("due date is required")
	}
	return nil
}

// Complete marks the action completed.
func (c *CAPA) Complete() {
	now := time.Now()
	c.CompletedAt = &now
	c.Status = StatusCompleted
}

// Verify registers verification signature credentials.
func (c *CAPA) Verify(verifier string) {
	now := time.Now()
	c.VerifiedBy = verifier
	c.VerifiedAt = &now
	c.Status = StatusVerified
}
