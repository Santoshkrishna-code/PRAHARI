package rehabilitation

import (
	"errors"
	"time"
)

// RehabilitationProgram tracks rehab status for injured workers.
type RehabilitationProgram struct {
	ID              string    `json:"id" db:"id"`
	HealthProfileID string    `json:"health_profile_id" db:"health_profile_id"`
	ProgramName     string    `json:"program_name" db:"program_name"`
	StartDate       time.Time `json:"start_date" db:"start_date"`
	EndDate         time.Time `json:"end_date" db:"end_date"`
	Status          string    `json:"status" db:"status"` // "PLANNED", "IN_PROGRESS", "COMPLETED", "FAILED"
	Notes           string    `json:"notes" db:"notes"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks rehab attributes.
func (r *RehabilitationProgram) Validate() error {
	if r.HealthProfileID == "" {
		return errors.New("health profile reference is required")
	}
	if r.ProgramName == "" {
		return errors.New("rehabilitation program name is required")
	}
	return nil
}
