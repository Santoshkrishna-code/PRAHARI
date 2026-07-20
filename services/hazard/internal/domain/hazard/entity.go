package hazard

import (
	"errors"
	"time"
)

// Hazard is the central aggregate root of the proactive hazard safety engine.
type Hazard struct {
	ID                string    `json:"id" db:"id"`
	HazardNumber      string    `json:"hazard_number" db:"hazard_number"`
	AssetID           string    `json:"asset_id,omitempty" db:"asset_id"`
	ContractorID      string    `json:"contractor_id,omitempty" db:"contractor_id"`
	HazardType        string    `json:"hazard_type" db:"hazard_type"`
	InitialRiskScore  int       `json:"initial_risk_score" db:"initial_risk_score"`
	ResidualRiskScore int       `json:"residual_risk_score" db:"residual_risk_score"`
	StatusCode        string    `json:"status_code" db:"status_code"`
	DepartmentID      string    `json:"department_id" db:"department_id"`
	Title             string    `json:"title" db:"title"`
	Description       string    `json:"description" db:"description"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
	IsDeleted         bool      `json:"is_deleted" db:"is_deleted"`
}

// Validate checks domain invariants.
func (h *Hazard) Validate() error {
	if h.Title == "" {
		return errors.New("hazard title is required")
	}
	if len(h.Title) > 200 {
		return errors.New("hazard title must not exceed 200 characters")
	}
	if h.DepartmentID == "" {
		return errors.New("department ID is required")
	}
	return nil
}
