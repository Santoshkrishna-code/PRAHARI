package renewableenergy

import (
	"errors"
	"time"
)

// EnergyRecord registers clean energy production or buy.
type EnergyRecord struct {
	ID             string    `json:"id" db:"id"`
	PlantID        string    `json:"plant_id" db:"plant_id"`
	GenerationType string    `json:"generation_type" db:"generation_type"` // "SOLAR", "WIND", "HYDRO", "PPA"
	AmountKWh      float64   `json:"amount_kwh" db:"amount_kwh"`
	CarbonOffsetKg float64   `json:"carbon_offset_kg" db:"carbon_offset_kg"`
	RecordedAt     time.Time `json:"recorded_at" db:"recorded_at"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

// Validate checks fields.
func (e *EnergyRecord) Validate() error {
	if e.PlantID == "" {
		return errors.New("plant ID reference is required")
	}
	if e.AmountKWh <= 0.0 {
		return errors.New("generation amount KWh must be positive")
	}
	return nil
}
