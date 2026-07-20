package esgobjective

import (
	"errors"
	"time"
)

// Objective structures corporate ESG target parameters.
type Objective struct {
	ID             string    `json:"id" db:"id"`
	BusinessUnitID string    `json:"business_unit_id" db:"business_unit_id"`
	Title          string    `json:"title" db:"title"` // e.g. "Net Zero Scope 1 by 2030"
	Category       string    `json:"category" db:"category"` // "ENVIRONMENTAL", "SOCIAL", "GOVERNANCE"
	TargetValue    float64   `json:"target_value" db:"target_value"`
	CurrentValue   float64   `json:"current_value" db:"current_value"`
	UnitOfMeasure  string    `json:"unit_of_measure" db:"unit_of_measure"`
	Deadline       time.Time `json:"deadline" db:"deadline"`
	Status         string    `json:"status" db:"status"` // "OBJECTIVE_DEFINED", "PUBLISHED", "ACHIEVED"
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks targets.
func (o *Objective) Validate() error {
	if o.BusinessUnitID == "" {
		return errors.New("business unit reference is required")
	}
	if o.Title == "" {
		return errors.New("objective title description is required")
	}
	return nil
}
