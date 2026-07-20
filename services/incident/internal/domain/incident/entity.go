package incident

import (
	"errors"
	"fmt"
	"time"
)

// Type classifies the nature of the reported event.
type Type string

const (
	TypeIncident        Type = "INCIDENT"
	TypeNearMiss        Type = "NEAR_MISS"
	TypeUnsafeAct       Type = "UNSAFE_ACT"
	TypeUnsafeCondition Type = "UNSAFE_CONDITION"
	TypeHazard          Type = "HAZARD"
)

// ValidTypes enumerates all accepted incident type classifications.
var ValidTypes = []Type{
	TypeIncident,
	TypeNearMiss,
	TypeUnsafeAct,
	TypeUnsafeCondition,
	TypeHazard,
}

// Incident is the core aggregate root of the Incident Management bounded context.
// It represents a single safety event reported within the organization, carrying
// the full lifecycle from draft through investigation to closure.
type Incident struct {
	ID             string    `json:"id" db:"id"`
	IncidentNumber string    `json:"incident_number" db:"incident_number"`
	Title          string    `json:"title" db:"title"`
	Description    string    `json:"description" db:"description"`
	Type           Type      `json:"type" db:"type"`
	CategoryID     string    `json:"category_id" db:"category_id"`
	SeverityLevel  string    `json:"severity_level" db:"severity_level"`
	PriorityLevel  string    `json:"priority_level" db:"priority_level"`
	StatusCode     string    `json:"status_code" db:"status_code"`
	ReporterID     string    `json:"reporter_id" db:"reporter_id"`
	AssigneeID     string    `json:"assignee_id,omitempty" db:"assignee_id"`
	DepartmentID   string    `json:"department_id" db:"department_id"`
	LocationID     string    `json:"location_id" db:"location_id"`
	LocationDetail string    `json:"location_detail,omitempty" db:"location_detail"`
	OccurredAt     time.Time `json:"occurred_at" db:"occurred_at"`
	ReportedAt     time.Time `json:"reported_at" db:"reported_at"`
	ResolvedAt     *time.Time `json:"resolved_at,omitempty" db:"resolved_at"`
	ClosedAt       *time.Time `json:"closed_at,omitempty" db:"closed_at"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	IsDeleted      bool      `json:"is_deleted" db:"is_deleted"`
}

// Validate enforces domain invariants on the incident aggregate.
func (i *Incident) Validate() error {
	if i.Title == "" {
		return errors.New("incident title is required")
	}
	if len(i.Title) > 500 {
		return errors.New("incident title must not exceed 500 characters")
	}
	if i.Description == "" {
		return errors.New("incident description is required")
	}
	if !i.isValidType() {
		return fmt.Errorf("invalid incident type: %s", i.Type)
	}
	if i.ReporterID == "" {
		return errors.New("reporter ID is required")
	}
	if i.DepartmentID == "" {
		return errors.New("department ID is required")
	}
	if i.OccurredAt.IsZero() {
		return errors.New("incident occurrence timestamp is required")
	}
	if i.OccurredAt.After(time.Now()) {
		return errors.New("incident occurrence timestamp cannot be in the future")
	}
	return nil
}

// isValidType checks whether the incident type is among accepted classifications.
func (i *Incident) isValidType() bool {
	for _, t := range ValidTypes {
		if i.Type == t {
			return true
		}
	}
	return false
}

// CanBeAssigned checks whether the incident is in a state that allows assignment.
func (i *Incident) CanBeAssigned() bool {
	return i.StatusCode == "UNDER_REVIEW" || i.StatusCode == "ASSIGNED"
}

// CanBeInvestigated checks whether the incident is in a state that allows investigation.
func (i *Incident) CanBeInvestigated() bool {
	return i.StatusCode == "ASSIGNED"
}

// CanBeResolved checks whether the incident lifecycle permits resolution.
func (i *Incident) CanBeResolved() bool {
	return i.StatusCode == "CAPA_IN_PROGRESS" || i.StatusCode == "INVESTIGATING"
}

// CanBeClosed checks whether the incident can be closed by a safety officer.
func (i *Incident) CanBeClosed() bool {
	return i.StatusCode == "RESOLVED"
}

// MarkResolved sets the resolution timestamp and transitions status.
func (i *Incident) MarkResolved() {
	now := time.Now()
	i.ResolvedAt = &now
	i.StatusCode = "RESOLVED"
	i.UpdatedAt = now
}

// MarkClosed sets the closure timestamp and transitions status.
func (i *Incident) MarkClosed() {
	now := time.Now()
	i.ClosedAt = &now
	i.StatusCode = "CLOSED"
	i.UpdatedAt = now
}
