package medicalexamination

import (
	"errors"
	"time"
)

// MedicalExamination holds scheduled or completed occupational exam metrics.
type MedicalExamination struct {
	ID              string    `json:"id" db:"id"`
	HealthProfileID string    `json:"health_profile_id" db:"health_profile_id"`
	ExamType        string    `json:"exam_type" db:"exam_type"` // "PRE_EMPLOYMENT", "PERIODIC", "EXIT", "FITNESS_FOR_DUTY"
	ExamDate        time.Time `json:"exam_date" db:"exam_date"`
	PhysicianID     string    `json:"physician_id" db:"physician_id"`
	ClinicID        string    `json:"clinic_id" db:"clinic_id"`
	VitalsBP        string    `json:"vitals_bp" db:"vitals_bp"`         // e.g. "120/80"
	VitalsPulse     int       `json:"vitals_pulse" db:"vitals_pulse"`   // bpm
	WeightKg        float64   `json:"weight_kg" db:"weight_kg"`
	HeightCm        float64   `json:"height_cm" db:"height_cm"`
	Findings        string    `json:"findings" db:"findings"`
	Recommendations string    `json:"recommendations" db:"recommendations"`
	OutcomeStatus   string    `json:"outcome_status" db:"outcome_status"` // "COMPLETED", "INCOMPLETE", "REFERRAL"
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks exam constraints.
func (e *MedicalExamination) Validate() error {
	if e.HealthProfileID == "" {
		return errors.New("health profile reference is required")
	}
	if e.ExamType == "" {
		return errors.New("exam type is required")
	}
	if e.PhysicianID == "" {
		return errors.New("physician ID is required")
	}
	return nil
}
