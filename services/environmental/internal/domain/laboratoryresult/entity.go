package laboratoryresult

import (
	"errors"
	"time"
)

// LaboratoryResult registers lab parameters evaluation.
type LaboratoryResult struct {
	ID             string    `json:"id" db:"id"`
	SampleID       string    `json:"sample_id" db:"sample_id"`
	LaboratoryID   string    `json:"laboratory_id" db:"laboratory_id"`
	AnalyteName    string    `json:"analyte_name" db:"analyte_name"` // "COD", "BOD", "LEAD", "BENZENE"
	AnalyteValue   float64   `json:"analyte_value" db:"analyte_value"`
	UnitOfMeasure  string    `json:"unit_of_measure" db:"unit_of_measure"`
	RegulatoryLimit float64  `json:"regulatory_limit" db:"regulatory_limit"`
	IsAbnormal     bool      `json:"is_abnormal" db:"is_abnormal"`
	TestDate       time.Time `json:"test_date" db:"test_date"`
	PhysicianNotes string    `json:"physician_notes" db:"physician_notes"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks results.
func (r *LaboratoryResult) Validate() error {
	if r.SampleID == "" {
		return errors.New("sample ID reference is required")
	}
	if r.AnalyteName == "" {
		return errors.New("analyte name is required")
	}
	return nil
}
