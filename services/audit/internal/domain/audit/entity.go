package audit

import (
	"errors"
	"time"
)

// Audit represents the central aggregate root of the Audit Management domain.
type Audit struct {
	ID               string    `json:"id" db:"id"`
	AuditNumber      string    `json:"audit_number" db:"audit_number"`
	AssetID          string    `json:"asset_id,omitempty" db:"asset_id"`
	DepartmentID     string    `json:"department_id" db:"department_id"`
	ComplianceRating float64   `json:"compliance_rating" db:"compliance_rating"`
	StatusCode       string    `json:"status_code" db:"status_code"`
	Title            string    `json:"title" db:"title"`
	Description      string    `json:"description" db:"description"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
	IsDeleted        bool      `json:"is_deleted" db:"is_deleted"`
}

// Validate checks domain invariants.
func (a *Audit) Validate() error {
	if a.Title == "" {
		return errors.New("audit title is required")
	}
	if len(a.Title) > 200 {
		return errors.New("audit title must not exceed 200 characters")
	}
	if a.DepartmentID == "" {
		return errors.New("department ID is required")
	}
	return nil
}
