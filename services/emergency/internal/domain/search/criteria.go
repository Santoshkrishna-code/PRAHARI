package search

import "time"

// Criteria defines multi-dimensional search parameters for industrial emergencies.
type Criteria struct {
	OrganizationID string     `json:"organization_id,omitempty"`
	PlantID        string     `json:"plant_id,omitempty"`
	UnitID         string     `json:"unit_id,omitempty"`
	Category       string     `json:"category,omitempty"`
	Severity       string     `json:"severity,omitempty"`
	CommanderID    string     `json:"commander_id,omitempty"`
	Status         string     `json:"status,omitempty"`
	StartDate      *time.Time `json:"start_date,omitempty"`
	EndDate        *time.Time `json:"end_date,omitempty"`
	Query          string     `json:"query,omitempty"`
	Limit          int        `json:"limit,omitempty"`
	Offset         int        `json:"offset,omitempty"`
}
