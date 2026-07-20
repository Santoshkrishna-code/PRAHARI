package efficiency

import (
	"errors"
	"time"
)

// Program registers ISO 50001 energy targets.
type Program struct {
	ID             string    `json:"id" db:"id"`
	PlantID        string    `json:"plant_id" db:"plant_id"`
	Title          string    `json:"title" db:"title"` // e.g. "Compressor VFD Upgrades"
	BaselineKWh    float64   `json:"baseline_kwh" db:"baseline_kwh"`
	TargetSavedKWh float64   `json:"target_saved_kwh" db:"target_saved_kwh"`
	ActualSavedKWh float64   `json:"actual_saved_kwh" db:"actual_saved_kwh"`
	Status         string    `json:"status" db:"status"` // "PLANNED", "IN_PROGRESS", "COMPLETED"
	StartDate      time.Time `json:"start_date" db:"start_date"`
	EndDate        time.Time `json:"end_date" db:"end_date"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

// Validate checks fields.
func (p *Program) Validate() error {
	if p.PlantID == "" {
		return errors.New("plant ID reference is required")
	}
	if p.Title == "" {
		return errors.New("program title description is required")
	}
	return nil
}
