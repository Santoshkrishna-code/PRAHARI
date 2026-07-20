package search

import "time"

// Criteria defines parameters for multi-dimensional search.
type Criteria struct {
	PlantID        string     `json:"plant_id,omitempty"`
	PhysicalState  string     `json:"physical_state,omitempty"`
	IsRestricted   *bool      `json:"is_restricted,omitempty"`
	Status         string     `json:"status,omitempty"`
	ExpiryBefore   *time.Time `json:"expiry_before,omitempty"`
	Query          string     `json:"query,omitempty"`
	Limit          int        `json:"limit,omitempty"`
	Offset         int        `json:"offset,omitempty"`
}
