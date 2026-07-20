package search

import "time"

// Criteria defines parameters to filter digital twin data.
type Criteria struct {
	TwinID    string     `json:"twin_id,omitempty"`
	PlantID   string     `json:"plant_id,omitempty"`
	Query     string     `json:"query,omitempty"`
	StartTime *time.Time `json:"start_time,omitempty"`
	EndTime   *time.Time `json:"end_time,omitempty"`
	Limit     int        `json:"limit,omitempty"`
	Offset    int        `json:"offset,omitempty"`
}
