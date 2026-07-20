package search

import "time"

// Criteria defines multi-dimensional search parameters for LOTO certificates, plans, and audits.
type Criteria struct {
	PlantID      string     `json:"plant_id,omitempty"`
	PlanID       string     `json:"plan_id,omitempty"`
	Status       string     `json:"status,omitempty"`
	PermitID     string     `json:"permit_id,omitempty"`
	StartDate    *time.Time `json:"start_date,omitempty"`
	EndDate      *time.Time `json:"end_date,omitempty"`
	Query        string     `json:"query,omitempty"`
	Limit        int        `json:"limit,omitempty"`
	Offset       int        `json:"offset,omitempty"`
}
