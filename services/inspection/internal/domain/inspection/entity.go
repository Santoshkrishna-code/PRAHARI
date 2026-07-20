package inspection

import (
	"errors"
	"fmt"
	"time"
)

// Type classifies safety audit inspections categories.
type Type string

const (
	TypeRoutine       Type = "ROUTINE"
	TypeSafety        Type = "SAFETY"
	TypePermit        Type = "PERMIT"
	TypeEquipment     Type = "EQUIPMENT"
	TypeEnvironmental Type = "ENVIRONMENTAL"
	TypeAudit         Type = "AUDIT"
)

// Inspection is the central aggregate root of the Inspection Management domain.
type Inspection struct {
	ID               string     `json:"id" db:"id"`
	InspectionNumber string     `json:"inspection_number" db:"inspection_number"`
	Title            string     `json:"title" db:"title"`
	Description      string     `json:"description" db:"description"`
	InspectionType   Type       `json:"inspection_type" db:"inspection_type"`
	StatusCode       string     `json:"status_code" db:"status_code"`
	ScheduleID       string     `json:"schedule_id,omitempty" db:"schedule_id"`
	InspectorID      string     `json:"inspector_id" db:"inspector_id"`
	DepartmentID     string     `json:"department_id" db:"department_id"`
	AssetID          string     `json:"asset_id,omitempty" db:"asset_id"`
	LinkedPermitID   string     `json:"linked_permit_id,omitempty" db:"linked_permit_id"`
	LinkedIncidentID string     `json:"linked_incident_id,omitempty" db:"linked_incident_id"`
	ComplianceScore  float64    `json:"compliance_score" db:"compliance_score"`
	StartedAt        *time.Time `json:"started_at,omitempty" db:"started_at"`
	CompletedAt      *time.Time `json:"completed_at,omitempty" db:"completed_at"`
	CreatedAt        time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at" db:"updated_at"`
	IsDeleted        bool       `json:"is_deleted" db:"is_deleted"`
}

// Validate checks domain invariants for Inspection.
func (i *Inspection) Validate() error {
	if i.Title == "" {
		return errors.New("inspection title is required")
	}
	if len(i.Title) > 200 {
		return errors.New("inspection title must not exceed 200 characters")
	}
	if i.InspectorID == "" {
		return errors.New("inspector ID is required")
	}
	if i.DepartmentID == "" {
		return errors.New("department ID is required")
	}
	return nil
}

// CanStart checks if state transitions are permitted.
func (i *Inspection) CanStart() bool {
	return i.StatusCode == "ASSIGNED" || i.StatusCode == "SCHEDULED"
}

// CanComplete checks execution completeness.
func (i *Inspection) CanComplete() bool {
	return i.StatusCode == "IN_PROGRESS"
}
