package compliance

import (
	"errors"
	"time"
)

// Compliance represents the central aggregate root of the Compliance Register.
type Compliance struct {
	ID               string    `json:"id" db:"id"`
	ComplianceNumber string    `json:"compliance_number" db:"compliance_number"`
	AssetID          string    `json:"asset_id,omitempty" db:"asset_id"`
	DepartmentID     string    `json:"department_id" db:"department_id"`
	ComplianceScore  float64   `json:"compliance_score" db:"compliance_score"`
	StatusCode       string    `json:"status_code" db:"status_code"`
	Title            string    `json:"title" db:"title"`
	Description      string    `json:"description" db:"description"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
	IsDeleted        bool      `json:"is_deleted" db:"is_deleted"`
}

// Validate checks domain invariants.
func (c *Compliance) Validate() error {
	if c.Title == "" {
		return errors.New("compliance register title is required")
	}
	if len(c.Title) > 200 {
		return errors.New("compliance register title must not exceed 200 characters")
	}
	if c.DepartmentID == "" {
		return errors.New("department ID is required")
	}
	return nil
}
