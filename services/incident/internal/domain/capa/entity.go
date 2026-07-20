package capa

import (
	"errors"
	"fmt"
	"time"
)

// Type classifies the CAPA as corrective or preventive.
type Type string

const (
	TypeCorrective Type = "CORRECTIVE"
	TypePreventive Type = "PREVENTIVE"
)

// Status defines the CAPA tracking lifecycle.
type Status string

const (
	StatusOpen       Status = "OPEN"
	StatusInProgress Status = "IN_PROGRESS"
	StatusCompleted  Status = "COMPLETED"
	StatusOverdue    Status = "OVERDUE"
	StatusVerified   Status = "VERIFIED"
)

// CAPA represents a Corrective and Preventive Action linked to an incident.
type CAPA struct {
	ID                  string     `json:"id" db:"id"`
	IncidentID          string     `json:"incident_id" db:"incident_id"`
	Type                Type       `json:"type" db:"type"`
	Description         string     `json:"description" db:"description"`
	AssigneeID          string     `json:"assignee_id" db:"assignee_id"`
	DueDate             time.Time  `json:"due_date" db:"due_date"`
	CompletedAt         *time.Time `json:"completed_at,omitempty" db:"completed_at"`
	Status              Status     `json:"status" db:"status"`
	VerifiedBy          string     `json:"verified_by,omitempty" db:"verified_by"`
	VerifiedAt          *time.Time `json:"verified_at,omitempty" db:"verified_at"`
	EffectivenessReview string     `json:"effectiveness_review,omitempty" db:"effectiveness_review"`
	CreatedAt           time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at" db:"updated_at"`
}

// Validate enforces domain invariants on the CAPA aggregate.
func (c *CAPA) Validate() error {
	if c.IncidentID == "" {
		return errors.New("incident ID is required for CAPA")
	}
	if c.Type != TypeCorrective && c.Type != TypePreventive {
		return fmt.Errorf("invalid CAPA type: %s", c.Type)
	}
	if c.Description == "" {
		return errors.New("CAPA description is required")
	}
	if c.AssigneeID == "" {
		return errors.New("CAPA assignee is required")
	}
	if c.DueDate.IsZero() {
		return errors.New("CAPA due date is required")
	}
	return nil
}

// Complete marks the CAPA as completed.
func (c *CAPA) Complete() {
	now := time.Now()
	c.CompletedAt = &now
	c.Status = StatusCompleted
	c.UpdatedAt = now
}

// Verify marks the CAPA as verified after effectiveness review.
func (c *CAPA) Verify(verifierID, review string) {
	now := time.Now()
	c.VerifiedBy = verifierID
	c.VerifiedAt = &now
	c.EffectivenessReview = review
	c.Status = StatusVerified
	c.UpdatedAt = now
}

// IsOverdue returns true if the CAPA has passed its due date without completion.
func (c *CAPA) IsOverdue() bool {
	return c.Status != StatusCompleted && c.Status != StatusVerified && time.Now().After(c.DueDate)
}

// CheckAndMarkOverdue transitions the CAPA to overdue status if applicable.
func (c *CAPA) CheckAndMarkOverdue() bool {
	if c.IsOverdue() && c.Status != StatusOverdue {
		c.Status = StatusOverdue
		c.UpdatedAt = time.Now()
		return true
	}
	return false
}
