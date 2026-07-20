package medicalclearance

import (
	"errors"
	"time"
)

// MedicalClearance validates worker eligibility for permits or onboarding.
type MedicalClearance struct {
	ID              string    `json:"id" db:"id"`
	HealthProfileID string    `json:"health_profile_id" db:"health_profile_id"`
	ClearanceDate   time.Time `json:"clearance_date" db:"clearance_date"`
	ExpiryDate      time.Time `json:"expiry_date" db:"expiry_date"`
	IsApproved      bool      `json:"is_approved" db:"is_approved"`
	ApprovedByID    string    `json:"approved_by_id" db:"approved_by_id"`
	ScopeOfWork     string    `json:"scope_of_work" db:"scope_of_work"`
	Notes           string    `json:"notes" db:"notes"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks clearance parameters.
func (c *MedicalClearance) Validate() error {
	if c.HealthProfileID == "" {
		return errors.New("health profile reference is required")
	}
	if c.ExpiryDate.Before(c.ClearanceDate) {
		return errors.New("expiry date must be after clearance date")
	}
	return nil
}
