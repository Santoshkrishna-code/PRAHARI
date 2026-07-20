package laboratoryresult

import (
	"errors"
	"time"
)

// LaboratoryResult maps medical panel/tests values.
type LaboratoryResult struct {
	ID              string    `json:"id" db:"id"`
	ExamID          string    `json:"exam_id" db:"exam_id"` // Optional link to parent MedicalExamination
	LaboratoryID    string    `json:"laboratory_id" db:"laboratory_id"`
	TestName        string    `json:"test_name" db:"test_name"` // "BLOOD_PANEL", "URINALYSIS", "DRUG_SCREEN"
	TestValue       string    `json:"test_value" db:"test_value"`
	ReferenceRange  string    `json:"reference_range" db:"reference_range"` // Normal range helper
	IsAbnormal      bool      `json:"is_abnormal" db:"is_abnormal"`
	TestDate        time.Time `json:"test_date" db:"test_date"`
	PhysicianNotes  string    `json:"physician_notes" db:"physician_notes"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks results fields.
func (r *LaboratoryResult) Validate() error {
	if r.LaboratoryID == "" {
		return errors.New("laboratory reference is required")
	}
	if r.TestName == "" {
		return errors.New("test name is required")
	}
	return nil
}
