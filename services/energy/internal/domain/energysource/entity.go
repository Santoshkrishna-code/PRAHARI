package energysource

import (
	"errors"
	"time"
)

// Source represents an industrial energy resource definition.
type Source struct {
	ID         string    `json:"id" db:"id"`
	SourceName string    `json:"source_name" db:"source_name"` // "ELECTRICITY", "DIESEL", "NATURAL_GAS"
	EnergyType string    `json:"energy_type" db:"energy_type"` // "RENEWABLE", "NON_RENEWABLE"
	GridRegion string    `json:"grid_region" db:"grid_region"`
	IsActive   bool      `json:"is_active" db:"is_active"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

// Validate checks standard values.
func (s *Source) Validate() error {
	if s.SourceName == "" {
		return errors.New("energy source name key is required")
	}
	return nil
}
