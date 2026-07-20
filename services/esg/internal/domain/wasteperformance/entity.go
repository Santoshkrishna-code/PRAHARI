package wasteperformance

import (
	"errors"
	"time"
)

// PerformanceRecord monitors overall waste objectives metrics.
type PerformanceRecord struct {
	ID             string    `json:"id" db:"id"`
	PlantID        string    `json:"plant_id" db:"plant_id"`
	HazardousKg    float64   `json:"hazardous_kg" db:"hazardous_kg"`
	NonHazardousKg float64   `json:"non_hazardous_kg" db:"non_hazardous_kg"`
	TotalWasteKg   float64   `json:"total_waste_kg" db:"total_waste_kg"`
	RecordedPeriod string    `json:"recorded_period" db:"recorded_period"`
	RecordedAt     time.Time `json:"recorded_at" db:"recorded_at"`
}

// Validate checks fields.
func (w *PerformanceRecord) Validate() error {
	if w.PlantID == "" {
		return errors.New("plant ID reference is required")
	}
	return nil
}

// EvaluateTotal sum values.
func (w *PerformanceRecord) EvaluateTotal() {
	w.TotalWasteKg = w.HazardousKg + w.NonHazardousKg
}
