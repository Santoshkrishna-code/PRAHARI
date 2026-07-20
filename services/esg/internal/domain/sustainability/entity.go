package sustainability

import (
	"errors"
	"time"
)

// Profile registers main ESG Business Unit profiles.
type Profile struct {
	ID             string    `json:"id" db:"id"`
	BusinessUnitID string    `json:"business_unit_id" db:"business_unit_id"`
	PlantID        string    `json:"plant_id" db:"plant_id"`
	DepartmentID   string    `json:"department_id" db:"department_id"`
	FTECount       int       `json:"fte_count" db:"fte_count"`
	RevenueUSD     float64   `json:"revenue_usd" db:"revenue_usd"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks profile values.
func (p *Profile) Validate() error {
	if p.BusinessUnitID == "" {
		return errors.New("business unit ID reference is required")
	}
	if p.PlantID == "" {
		return errors.New("plant ID reference is required")
	}
	return nil
}
