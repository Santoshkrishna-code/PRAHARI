package illness

import (
	"errors"
	"time"
)

// OccupationalIllness registers job-related disease or health deterioration diagnoses.
type OccupationalIllness struct {
	ID              string    `json:"id" db:"id"`
	HealthProfileID string    `json:"health_profile_id" db:"health_profile_id"`
	IncidentID      string    `json:"incident_id" db:"incident_id"` // Optional link to reactive safety incident
	IllnessName     string    `json:"illness_name" db:"illness_name"`
	ICD10Code       string    `json:"icd10_code" db:"icd10_code"`
	DiagnosisDate   time.Time `json:"diagnosis_date" db:"diagnosis_date"`
	SeverityCode    string    `json:"severity_code" db:"severity_code"` // "MILD", "MODERATE", "SEVERE"
	Status          string    `json:"status" db:"status"`               // "ACTIVE", "RECOVERED", "REHABILITATION"
	Notes           string    `json:"notes" db:"notes"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks illness details.
func (i *OccupationalIllness) Validate() error {
	if i.HealthProfileID == "" {
		return errors.New("health profile reference is required")
	}
	if i.IllnessName == "" {
		return errors.New("illness name is required")
	}
	return nil
}
