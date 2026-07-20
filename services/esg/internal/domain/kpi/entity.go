package kpi

import (
	"errors"
	"time"
)

// KPI represents corporate ESG indicator targets.
type KPI struct {
	ID            string    `json:"id" db:"id"`
	Code          string    `json:"code" db:"code"` // e.g. "ENERGY_INTENSITY"
	Title         string    `json:"title" db:"title"`
	TargetValue   float64   `json:"target_value" db:"target_value"`
	CurrentValue  float64   `json:"current_value" db:"current_value"`
	UnitOfMeasure string    `json:"unit_of_measure" db:"unit_of_measure"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

// Validate checks KPI.
func (k *KPI) Validate() error {
	if k.Code == "" {
		return errors.New("kpi unique code key is required")
	}
	return nil
}
