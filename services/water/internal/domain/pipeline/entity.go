package pipeline

import "time"

// Pipeline represents a water distribution pipeline segment.
type Pipeline struct {
	ID             string    `json:"id"`
	PlantID        string    `json:"plant_id"`
	PipelineName   string    `json:"pipeline_name"`
	DiameterMM     float64   `json:"diameter_mm"`
	LengthMeters   float64   `json:"length_meters"`
	Material       string    `json:"material"`
	PressureBarMax float64   `json:"pressure_bar_max"`
	FromNode       string    `json:"from_node"`
	ToNode         string    `json:"to_node"`
	AssetID        string    `json:"asset_id,omitempty"`
	IsActive       bool      `json:"is_active"`
	CreatedAt      time.Time `json:"created_at"`
}
