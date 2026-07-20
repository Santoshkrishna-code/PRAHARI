package search

import "time"

// Criteria defines multi-dimensional search parameters for shifts, logbooks, and handovers.
type Criteria struct {
	PlantID      string     `json:"plant_id,omitempty"`
	UnitID       string     `json:"unit_id,omitempty"`
	SupervisorID string     `json:"supervisor_id,omitempty"`
	Status       string     `json:"status,omitempty"`
	StartDate    *time.Time `json:"start_date,omitempty"`
	EndDate      *time.Time `json:"end_date,omitempty"`
	Query        string     `json:"query,omitempty"`
	Limit        int        `json:"limit,omitempty"`
	Offset       int        `json:"offset,omitempty"`
}
