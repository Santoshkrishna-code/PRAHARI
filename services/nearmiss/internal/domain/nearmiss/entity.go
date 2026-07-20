package nearmiss

import (
	"errors"
	"time"
)

// NearMiss is the central aggregate root of the Near Miss domain.
type NearMiss struct {
	ID             string    `json:"id" db:"id"`
	NearMissNumber string    `json:"near_miss_number" db:"near_miss_number"`
	AssetID        string    `json:"asset_id,omitempty" db:"asset_id"`
	ContractorID   string    `json:"contractor_id,omitempty" db:"contractor_id"`
	Classification string    `json:"classification" db:"classification"`
	SeverityLevel  string    `json:"severity_level" db:"severity_level"`
	StatusCode     string    `json:"status_code" db:"status_code"`
	DepartmentID   string    `json:"department_id" db:"department_id"`
	Title          string    `json:"title" db:"title"`
	Description    string    `json:"description" db:"description"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	IsDeleted      bool      `json:"is_deleted" db:"is_deleted"`
}

// Validate checks domain invariants.
func (nm *NearMiss) Validate() error {
	if nm.Title == "" {
		return errors.New("near miss title is required")
	}
	if len(nm.Title) > 200 {
		return errors.New("near miss title must not exceed 200 characters")
	}
	if nm.DepartmentID == "" {
		return errors.New("department ID is required")
	}
	return nil
}
