package search

import "time"

// Criteria defines multi-dimensional search parameters for PPE items, catalogs, and issuances.
type Criteria struct {
	PlantID        string     `json:"plant_id,omitempty"`
	PPEID          string     `json:"ppe_id,omitempty"`
	CategoryID     string     `json:"category_id,omitempty"`
	Status         string     `json:"status,omitempty"`
	IssuedToID     string     `json:"issued_to_id,omitempty"`
	ExpiryBefore   *time.Time `json:"expiry_before,omitempty"`
	Query          string     `json:"query,omitempty"`
	Limit          int        `json:"limit,omitempty"`
	Offset         int        `json:"offset,omitempty"`
}
