package surveillance

import (
	"errors"
	"time"
)

// HealthSurveillance tracks structured medical monitoring programs.
type HealthSurveillance struct {
	ID              string    `json:"id" db:"id"`
	HealthProfileID string    `json:"health_profile_id" db:"health_profile_id"`
	ProgramType     string    `json:"program_type" db:"program_type"` // e.g. "HEARING_CONSERVATION", "VISION_PROGRAM"
	StartDate       time.Time `json:"start_date" db:"start_date"`
	NextDueDate     time.Time `json:"next_due_date" db:"next_due_date"`
	Status          string    `json:"status" db:"status"` // "ACTIVE", "COMPLETED", "SUSPENDED"
	Notes           string    `json:"notes" db:"notes"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks program records.
func (s *HealthSurveillance) Validate() error {
	if s.HealthProfileID == "" {
		return errors.New("health profile reference is required")
	}
	if s.ProgramType == "" {
		return errors.New("program type is required")
	}
	return nil
}
