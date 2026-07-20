package risk

import (
	"errors"
	"time"
)

// Risk represents the central aggregate root of the enterprise risk register.
type Risk struct {
	ID                string    `json:"id" db:"id"`
	RiskNumber        string    `json:"risk_number" db:"risk_number"`
	AssetID           string    `json:"asset_id,omitempty" db:"asset_id"`
	DepartmentID      string    `json:"department_id" db:"department_id"`
	InherentRiskScore int       `json:"inherent_risk_score" db:"inherent_risk_score"`
	ResidualRiskScore int       `json:"residual_risk_score" db:"residual_risk_score"`
	StatusCode        string    `json:"status_code" db:"status_code"`
	Title             string    `json:"title" db:"title"`
	Description       string    `json:"description" db:"description"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
	IsDeleted         bool      `json:"is_deleted" db:"is_deleted"`
}

// Validate checks domain invariants.
func (r *Risk) Validate() error {
	if r.Title == "" {
		return errors.New("risk register title is required")
	}
	if len(r.Title) > 200 {
		return errors.New("risk register title must not exceed 200 characters")
	}
	if r.DepartmentID == "" {
		return errors.New("department ID is required")
	}
	return nil
}
