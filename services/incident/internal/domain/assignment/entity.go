package assignment

import (
	"errors"
	"fmt"
	"time"
)

// Role defines the capacity in which a user is assigned to an incident.
type Role string

const (
	RoleInvestigator  Role = "INVESTIGATOR"
	RoleSafetyOfficer Role = "SAFETY_OFFICER"
	RoleManager       Role = "MANAGER"
	RoleSupervisor    Role = "SUPERVISOR"
)

// ValidRoles enumerates all accepted assignment roles.
var ValidRoles = []Role{
	RoleInvestigator,
	RoleSafetyOfficer,
	RoleManager,
	RoleSupervisor,
}

// Assignment represents the binding of a user to an incident in a specific role.
type Assignment struct {
	ID         string     `json:"id" db:"id"`
	IncidentID string     `json:"incident_id" db:"incident_id"`
	AssigneeID string     `json:"assignee_id" db:"assignee_id"`
	AssignerID string     `json:"assigner_id" db:"assigner_id"`
	Role       Role       `json:"role" db:"role"`
	AssignedAt time.Time  `json:"assigned_at" db:"assigned_at"`
	AcceptedAt *time.Time `json:"accepted_at,omitempty" db:"accepted_at"`
	Note       string     `json:"note,omitempty" db:"note"`
	IsActive   bool       `json:"is_active" db:"is_active"`
}

// Validate enforces domain invariants on the assignment aggregate.
func (a *Assignment) Validate() error {
	if a.IncidentID == "" {
		return errors.New("incident ID is required for assignment")
	}
	if a.AssigneeID == "" {
		return errors.New("assignee ID is required for assignment")
	}
	if a.AssignerID == "" {
		return errors.New("assigner ID is required for assignment")
	}
	if !a.isValidRole() {
		return fmt.Errorf("invalid assignment role: %s", a.Role)
	}
	return nil
}

// Accept records that the assignee has acknowledged the assignment.
func (a *Assignment) Accept() {
	now := time.Now()
	a.AcceptedAt = &now
}

// IsAccepted returns true if the assignee has acknowledged the assignment.
func (a *Assignment) IsAccepted() bool {
	return a.AcceptedAt != nil
}

// isValidRole checks whether the role is among accepted assignment roles.
func (a *Assignment) isValidRole() bool {
	for _, r := range ValidRoles {
		if a.Role == r {
			return true
		}
	}
	return false
}
