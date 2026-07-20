package search

// Criteria defines multi-dimensional search parameters for integration connectors and jobs.
type Criteria struct {
	PlantID string `json:"plant_id,omitempty"`
	Query   string `json:"query,omitempty"`
	Limit   int    `json:"limit,omitempty"`
	Offset  int    `json:"offset,omitempty"`
}
