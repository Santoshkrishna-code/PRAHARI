package finding

import (
	"errors"
	"time"
)

// Severity defines safety threat categories.
type Severity string

const (
	SeverityMinor        Severity = "MINOR"
	SeverityMajor        Severity = "MAJOR"
	SeverityCritical     Severity = "CRITICAL"
	SeverityCatastrophic Severity = "CATASTROPHIC"
)

// Priority defines scheduling response targets.
type Priority string

const (
	PriorityLow       Priority = "LOW"
	PriorityMedium    Priority = "MEDIUM"
	PriorityHigh      Priority = "HIGH"
	PriorityEmergency Priority = "EMERGENCY"
)

// Status defines tracking lifecycle stages.
type Status string

const (
	StatusOpen       Status = "OPEN"
	StatusAssigned   Status = "ASSIGNED"
	StatusInProgress Status = "IN_PROGRESS"
	StatusResolved   Status = "RESOLVED"
	StatusVerified   Status = "VERIFIED"
	StatusClosed     Status = "CLOSED"
)

// Finding represents a critical issue identified during walkthrough inspections.
type Finding struct {
	ID             string    `json:"id" db:"id"`
	InspectionID   string    `json:"inspection_id" db:"inspection_id"`
	ChecklistItemID string   `json:"checklist_item_id" db:"checklist_item_id"`
	Description    string    `json:"description" db:"description"`
	Severity       Severity  `json:"severity" db:"severity"`
	Priority       Priority  `json:"priority" db:"priority"`
	Status         Status    `json:"status" db:"status"`
	IdentifiedBy   string    `json:"identified_by" db:"identified_by"`
	IdentifiedAt   time.Time `json:"identified_at" db:"identified_at"`
}

// Validate checks domain invariants for Finding.
func (f *Finding) Validate() error {
	if f.InspectionID == "" {
		return errors.New("inspection ID is required for finding")
	}
	if f.ChecklistItemID == "" {
		return errors.New("checklist item ID is required")
	}
	if f.Description == "" {
		return errors.New("finding description is required")
	}
	if f.IdentifiedBy == "" {
		return errors.New("identifier identity is required")
	}
	return nil
}

// IsCritical returns true if severity warrants auto-incident escalation.
func (f *Finding) IsCritical() bool {
	return f.Severity == SeverityCritical || f.Severity == SeverityCatastrophic
}
