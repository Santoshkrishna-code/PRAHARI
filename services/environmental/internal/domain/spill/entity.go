package spill

import (
	"errors"
	"time"
)

// ChemicalSpill registers accidental releases of chemicals or hydrocarbons.
type ChemicalSpill struct {
	ID             string    `json:"id" db:"id"`
	PlantID        string    `json:"plant_id" db:"plant_id"`
	IncidentID     string    `json:"incident_id" db:"incident_id"` // Link to reactive incident record
	ChemicalName   string    `json:"chemical_name" db:"chemical_name"`
	VolumeSpilled  float64   `json:"volume_spilled" db:"volume_spilled"`
	UnitOfMeasure  string    `json:"unit_of_measure" db:"unit_of_measure"` // "LITERS", "GALLONS"
	Contained      bool      `json:"contained" db:"contained"`
	ReachedWater   bool      `json:"reached_water" db:"reached_water"`
	ResponseAction string    `json:"response_action" db:"response_action"`
	SpillTime      time.Time `json:"spill_time" db:"spill_time"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

// Validate checks spill fields.
func (s *ChemicalSpill) Validate() error {
	if s.PlantID == "" {
		return errors.New("plant ID is required")
	}
	if s.ChemicalName == "" {
		return errors.New("spilled chemical name is required")
	}
	return nil
}
