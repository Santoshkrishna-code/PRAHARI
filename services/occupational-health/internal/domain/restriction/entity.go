package restriction

import (
	"errors"
	"time"
)

// MedicalRestriction maps limitations applied to a worker's scope of duty.
type MedicalRestriction struct {
	ID              string    `json:"id" db:"id"`
	HealthProfileID string    `json:"health_profile_id" db:"health_profile_id"`
	RestrictionCode string    `json:"restriction_code" db:"restriction_code"` // "NO_HEIGHT", "NO_HEAVY_LIFT", "NO_CHEMICAL"
	Description     string    `json:"description" db:"description"`
	StartDate       time.Time `json:"start_date" db:"start_date"`
	EndDate         time.Time `json:"end_date" db:"end_date"` // Optional for permanent restrictions
	IsPermanent     bool      `json:"is_permanent" db:"is_permanent"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks restriction dates.
func (r *MedicalRestriction) Validate() error {
	if r.HealthProfileID == "" {
		return errors.New("health profile reference is required")
	}
	if r.RestrictionCode == "" {
		return errors.New("restriction code is required")
	}
	if !r.IsPermanent && r.EndDate.Before(r.StartDate) {
		return errors.New("temporary restrictions must have a valid end date")
	}
	return nil
}
