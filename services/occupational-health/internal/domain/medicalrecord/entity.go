package medicalrecord

import (
	"errors"
	"time"
)

// MedicalRecord stores general health logs, notes, and clinical histories.
type MedicalRecord struct {
	ID              string    `json:"id" db:"id"`
	HealthProfileID string    `json:"health_profile_id" db:"health_profile_id"`
	RecordDate      time.Time `json:"record_date" db:"record_date"`
	RecordType      string    `json:"record_type" db:"record_type"` // "CLINICAL_NOTE", "CONSULTATION", "HISTORY"
	PhysicianID     string    `json:"physician_id" db:"physician_id"`
	DiagnosisCode   string    `json:"diagnosis_code" db:"diagnosis_code"` // ICD-10 standard codes
	ClinicalNotes   string    `json:"clinical_notes" db:"clinical_notes"` // Encrypted at DB layer in production
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks entity values.
func (r *MedicalRecord) Validate() error {
	if r.HealthProfileID == "" {
		return errors.New("health profile reference is required")
	}
	if r.RecordType == "" {
		return errors.New("record type is required")
	}
	return nil
}
