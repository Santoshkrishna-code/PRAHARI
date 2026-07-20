package emissionfactor

import (
	"errors"
	"time"
)

// Factor maps conversions.
type Factor struct {
	ID            string    `json:"id" db:"id"`
	SourceName    string    `json:"source_name" db:"source_name"` // "DIESEL", "GRID_ELECTRICITY"
	FactorValue   float64   `json:"factor_value" db:"factor_value"` // co2 kg per unit
	UnitOfMeasure string    `json:"unit_of_measure" db:"unit_of_measure"` // "LITER", "KWH"
	StandardName  string    `json:"standard_name" db:"standard_name"` // "GHG_PROTOCOL", "DEFRA"
	Year          int       `json:"year" db:"year"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}

// Validate checks factor.
func (f *Factor) Validate() error {
	if f.SourceName == "" {
		return errors.New("emission factor source name is required")
	}
	if f.FactorValue <= 0.0 {
		return errors.New("factor multiplier value must be positive")
	}
	return nil
}
