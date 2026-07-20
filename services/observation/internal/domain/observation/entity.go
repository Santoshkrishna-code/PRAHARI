package observation

import (
	"errors"
	"time"
)

// Observation is the central aggregate root of the Behavior-Based Safety domain.
type Observation struct {
	ID                string    `json:"id" db:"id"`
	ObservationNumber string    `json:"observation_number" db:"observation_number"`
	AssetID           string    `json:"asset_id,omitempty" db:"asset_id"`
	ContractorID      string    `json:"contractor_id,omitempty" db:"contractor_id"`
	ObservationType   string    `json:"observation_type" db:"observation_type"` // Safe/Unsafe Behavior/Condition
	StatusCode        string    `json:"status_code" db:"status_code"`
	DepartmentID      string    `json:"department_id" db:"department_id"`
	Title             string    `json:"title" db:"title"`
	Description       string    `json:"description" db:"description"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
	IsDeleted         bool      `json:"is_deleted" db:"is_deleted"`
}

// Validate checks domain invariants.
func (o *Observation) Validate() error {
	if o.Title == "" {
		return errors.New("observation title is required")
	}
	if len(o.Title) > 200 {
		return errors.New("observation title must not exceed 200 characters")
	}
	if o.DepartmentID == "" {
		return errors.New("department ID is required")
	}
	return nil
}
}
