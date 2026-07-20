package permit

import (
	"errors"
	"fmt"
	"time"
)

// RiskLevel defines the severity/class of hazards associated with a permit.
type RiskLevel string

const (
	RiskLevelLow      RiskLevel = "LOW"
	RiskLevelMedium   RiskLevel = "MEDIUM"
	RiskLevelHigh     RiskLevel = "HIGH"
	RiskLevelCritical RiskLevel = "CRITICAL"
)

// Permit is the central aggregate root of the Permit-to-Work domain.
// It manages the lifecycle and constraints of work authorizations.
type Permit struct {
	ID                string     `json:"id" db:"id"`
	PermitNumber      string     `json:"permit_number" db:"permit_number"`
	Title             string     `json:"title" db:"title"`
	Description       string     `json:"description" db:"description"`
	PermitTypeID      string     `json:"permit_type_id" db:"permit_type_id"`
	StatusCode        string     `json:"status_code" db:"status_code"`
	RiskLevel         RiskLevel  `json:"risk_level" db:"risk_level"`
	ApplicantID       string     `json:"applicant_id" db:"applicant_id"`
	SupervisorID      string     `json:"supervisor_id" db:"supervisor_id"`
	IssuerID          string     `json:"issuer_id" db:"issuer_id"`
	ReceiverID        string     `json:"receiver_id" db:"receiver_id"`
	DepartmentID      string     `json:"department_id" db:"department_id"`
	ContractorID      string     `json:"contractor_id,omitempty" db:"contractor_id"`
	WorkAreaID        string     `json:"work_area_id" db:"work_area_id"`
	WorkDescription   string     `json:"work_description" db:"work_description"`
	PlannedStartAt    time.Time  `json:"planned_start_at" db:"planned_start_at"`
	PlannedEndAt      time.Time  `json:"planned_end_at" db:"planned_end_at"`
	ActualStartAt     *time.Time `json:"actual_start_at,omitempty" db:"actual_start_at"`
	ActualEndAt       *time.Time `json:"actual_end_at,omitempty" db:"actual_end_at"`
	ValidUntil        *time.Time `json:"valid_until,omitempty" db:"valid_until"`
	ExtensionCount    int        `json:"extension_count" db:"extension_count"`
	LinkedIncidentID  string     `json:"linked_incident_id,omitempty" db:"linked_incident_id"`
	CreatedAt         time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at" db:"updated_at"`
	IsDeleted         bool       `json:"is_deleted" db:"is_deleted"`
}

// Validate checks domain invariants for a Permit.
func (p *Permit) Validate() error {
	if p.Title == "" {
		return errors.New("permit title is required")
	}
	if len(p.Title) > 200 {
		return errors.New("permit title must not exceed 200 characters")
	}
	if p.PermitTypeID == "" {
		return errors.New("permit type ID is required")
	}
	if p.ApplicantID == "" {
		return errors.New("applicant ID is required")
	}
	if p.DepartmentID == "" {
		return errors.New("department ID is required")
	}
	if p.WorkAreaID == "" {
		return errors.New("work area ID is required")
	}
	if p.PlannedStartAt.IsZero() {
		return errors.New("planned start date is required")
	}
	if p.PlannedEndAt.IsZero() {
		return errors.New("planned end date is required")
	}
	if p.PlannedEndAt.Before(p.PlannedStartAt) {
		return errors.New("planned end date must be after planned start date")
	}
	return nil
}

// IsExpired checks if the permit validity window has passed.
func (p *Permit) IsExpired() bool {
	if p.ValidUntil == nil {
		return false
	}
	return time.Now().After(*p.ValidUntil)
}

// CanBeExtended determines if the permit status and extension count permit extensions.
func (p *Permit) CanBeExtended() bool {
	return p.StatusCode == "ACTIVE" && p.ExtensionCount < 3 && !p.IsExpired()
}

// CanBeSuspended checks if the permit is currently in a state that permits suspension.
func (p *Permit) CanBeSuspended() bool {
	return p.StatusCode == "ACTIVE"
}

// CanBeActivated checks if the permit can transition to an active work state.
func (p *Permit) CanBeActivated() bool {
	return p.StatusCode == "ISSUED" || p.StatusCode == "SUSPENDED"
}
