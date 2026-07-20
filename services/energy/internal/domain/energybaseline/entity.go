package energybaseline

import (
	"errors"
	"time"
)

// Baseline tracks facility reference benchmarks.
type Baseline struct {
	ID             string    `json:"id" db:"id"`
	PlantID        string    `json:"plant_id" db:"plant_id"`
	BaselineYear   int       `json:"baseline_year" db:"baseline_year"`
	TotalKWh       float64   `json:"total_kwh" db:"total_kwh"`
	IntensityScore float64   `json:"intensity_score" db:"intensity_score"` // KWh per unit of production
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

// Validate checks targets.
func (b *Baseline) Validate() error {
	if b.PlantID == "" {
		return errors.New("plant ID reference is required")
	}
	return nil
}
