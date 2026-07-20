package search

import "time"

// Criteria defines multi-dimensional search parameters for analytics dashboards.
type Criteria struct {
	PlantID   string     `json:"plant_id,omitempty"`
	MetricKey string     `json:"metric_key,omitempty"`
	StartAt   *time.Time `json:"start_at,omitempty"`
	EndAt     *time.Time `json:"end_at,omitempty"`
	Limit     int        `json:"limit,omitempty"`
	Offset    int        `json:"offset,omitempty"`
}
