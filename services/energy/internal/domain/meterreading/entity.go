package meterreading

import (
	"errors"
	"time"
)

// Reading registers metrics.
type Reading struct {
	ID             string    `json:"id" db:"id"`
	MeterID        string    `json:"meter_id" db:"meter_id"`
	ReadingValue   float64   `json:"reading_value" db:"reading_value"`
	UnitOfMeasure  string    `json:"unit_of_measure" db:"unit_of_measure"` // "KWH", "M3", "KG_HR"
	Multiplier     float64   `json:"multiplier" db:"multiplier"`
	ActivePowerKW  float64   `json:"active_power_kw" db:"active_power_kw"`
	ReactivePowerVAR float64 `json:"reactive_power_var" db:"reactive_power_var"`
	PowerFactor    float64   `json:"power_factor" db:"power_factor"`
	ReadingTime    time.Time `json:"reading_time" db:"reading_time"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

// Validate checks reading limits.
func (r *Reading) Validate() error {
	if r.MeterID == "" {
		return errors.New("meter reference ID is required")
	}
	if r.ReadingValue < 0.0 {
		return errors.New("reading value cannot be negative")
	}
	return nil
}
