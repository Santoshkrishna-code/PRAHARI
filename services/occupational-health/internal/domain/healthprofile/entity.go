package healthprofile

import (
	"errors"
	"time"
)

// HealthProfile represents the aggregate root for a worker's health compliance state.
type HealthProfile struct {
	ID              string    `json:"id" db:"id"`
	WorkerID        string    `json:"worker_id" db:"worker_id"`
	WorkerType      string    `json:"worker_type" db:"worker_type"` // "EMPLOYEE" or "CONTRACTOR"
	DepartmentID    string    `json:"department_id" db:"department_id"`
	ClearanceStatus string    `json:"clearance_status" db:"clearance_status"` // e.g. "CLEARED", "RESTRICTED", "UNFIT"
	MedicalStatus   string    `json:"medical_status" db:"medical_status"`     // e.g. status.Code
	BloodType       string    `json:"blood_type" db:"blood_type"`
	DateOfBirth     time.Time `json:"date_of_birth" db:"date_of_birth"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
	IsDeleted       bool      `json:"is_deleted" db:"is_deleted"`
}

// Validate checks domain invariants.
func (p *HealthProfile) Validate() error {
	if p.WorkerID == "" {
		return errors.New("worker ID is required")
	}
	if p.WorkerType != "EMPLOYEE" && p.WorkerType != "CONTRACTOR" {
		return errors.New("worker type must be EMPLOYEE or CONTRACTOR")
	}
	if p.DepartmentID == "" {
		return errors.New("department ID is required")
	}
	return nil
}
