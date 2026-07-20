package energyefficiency

import (
	"errors"
	"time"
)

// EfficiencyRecord registers baseline vs actual targets.
type EfficiencyRecord struct {
	ID             string    `json:"id" db:"id"`
	PlantID        string    `json:"plant_id" db:"plant_id"`
	BaselineKWh    float64   `json:"baseline_kwh" db:"baseline_kwh"`
	ActualKWh      float64   `json:"actual_kwh" db:"actual_kwh"`
	EfficiencyRate float64   `json:"efficiency_rate" db:"efficiency_rate"` // (Baseline - Actual) / Baseline * 100
	RecordedPeriod string    `json:"recorded_period" db:"recorded_period"`  // e.g. "Q1-2026"
	RecordedAt     time.Time `json:"recorded_at" db:"recorded_at"`
}

// Validate checks fields.
func (e *EfficiencyRecord) Validate() error {
	if e.PlantID == "" {
		return errors.New("plant ID reference is required")
	}
	return nil
}

// EvaluateRate runs rate calculations.
func (e *EfficiencyRecord) EvaluateRate() {
	if e.BaselineKWh > 0.0 {
		e.EfficiencyRate = ((e.BaselineKWh - e.ActualKWh) / e.BaselineKWh) * 100.0
	}
}
