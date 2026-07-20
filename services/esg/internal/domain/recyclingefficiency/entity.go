package recyclingefficiency

import (
	"errors"
	"time"
)

// RecyclingKPI tracks trash conversion.
type RecyclingKPI struct {
	ID             string    `json:"id" db:"id"`
	PlantID        string    `json:"plant_id" db:"plant_id"`
	TotalWasteKg   float64   `json:"total_waste_kg" db:"total_waste_kg"`
	RecycledWasteKg float64  `json:"recycled_waste_kg" db:"recycled_waste_kg"`
	RecycleRate    float64   `json:"recycle_rate" db:"recycle_rate"` // Recycled / Total * 100
	RecordedAt     time.Time `json:"recorded_at" db:"recorded_at"`
}

// Validate checks limits.
func (r *RecyclingKPI) Validate() error {
	if r.PlantID == "" {
		return errors.New("plant ID reference is required")
	}
	return nil
}

// EvaluateRate runs rate checks.
func (r *RecyclingKPI) EvaluateRate() {
	if r.TotalWasteKg > 0.0 {
		r.RecycleRate = (r.RecycledWasteKg / r.TotalWasteKg) * 100.0
	}
}
