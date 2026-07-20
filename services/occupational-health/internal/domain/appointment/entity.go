package appointment

import (
	"errors"
	"time"
)

// Appointment tracks schedules for examinations.
type Appointment struct {
	ID              string    `json:"id" db:"id"`
	HealthProfileID string    `json:"health_profile_id" db:"health_profile_id"`
	ClinicID        string    `json:"clinic_id" db:"clinic_id"`
	PhysicianID     string    `json:"physician_id" db:"physician_id"`
	ScheduledTime   time.Time `json:"scheduled_time" db:"scheduled_time"`
	Purpose         string    `json:"purpose" db:"purpose"` // "PERIODIC_MEDICAL", "PRE_EMPLOYMENT"
	Status          string    `json:"status" db:"status"`   // "PENDING", "COMPLETED", "CANCELLED", "NOSHOW"
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks scheduled parameters.
func (a *Appointment) Validate() error {
	if a.HealthProfileID == "" {
		return errors.New("health profile reference is required")
	}
	if a.Purpose == "" {
		return errors.New("appointment purpose is required")
	}
	return nil
}
