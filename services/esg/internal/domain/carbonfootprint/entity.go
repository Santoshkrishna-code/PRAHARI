package carbonfootprint

import (
	"errors"
	"time"
)

// Footprint records overall intensities.
type Footprint struct {
	ID             string    `json:"id" db:"id"`
	BusinessUnitID string    `json:"business_unit_id" db:"business_unit_id"`
	ReportingYear  int       `json:"reporting_year" db:"reporting_year"`
	TotalEmissions float64   `json:"total_emissions" db:"total_emissions"` // in metric tonnes CO2e
	CarbonIntensity float64  `json:"carbon_intensity" db:"carbon_intensity"` // CO2e tonnes per FTE or USD
	CalculatedAt   time.Time `json:"calculated_at" db:"calculated_at"`
}

// Validate checks targets.
func (f *Footprint) Validate() error {
	if f.BusinessUnitID == "" {
		return errors.New("business unit reference is required")
	}
	return nil
}
