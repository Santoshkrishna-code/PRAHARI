package search

import "time"

// Criteria defines multi-dimensional search parameters for instruments, calibrations, and schedules.
type Criteria struct {
	PlantID        string     `json:"plant_id,omitempty"`
	InstrumentID   string     `json:"instrument_id,omitempty"`
	InstrumentType string     `json:"instrument_type,omitempty"`
	Status         string     `json:"status,omitempty"`
	ExpiryBefore   *time.Time `json:"expiry_before,omitempty"`
	Query          string     `json:"query,omitempty"`
	Limit          int        `json:"limit,omitempty"`
	Offset         int        `json:"offset,omitempty"`
}
