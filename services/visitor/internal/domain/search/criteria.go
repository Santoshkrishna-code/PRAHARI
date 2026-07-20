package search

import "time"

// Criteria defines multi-dimensional search parameters for visitors, visits, and gate passes.
type Criteria struct {
	OrganizationID string     `json:"organization_id,omitempty"`
	PlantID        string     `json:"plant_id,omitempty"`
	HostID         string     `json:"host_id,omitempty"`
	Contractor     string     `json:"contractor,omitempty"`
	VisitorCategory string    `json:"visitor_category,omitempty"`
	Vehicle        string     `json:"vehicle,omitempty"`
	Badge          string     `json:"badge,omitempty"`
	Status         string     `json:"status,omitempty"`
	StartDate      *time.Time `json:"start_date,omitempty"`
	EndDate        *time.Time `json:"end_date,omitempty"`
	Query          string     `json:"query,omitempty"`
	Limit          int        `json:"limit,omitempty"`
	Offset         int        `json:"offset,omitempty"`
}
