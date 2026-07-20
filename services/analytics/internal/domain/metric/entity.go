package metric

import "time"

// Metric represents an aggregated count or telemetry parameter record.
type Metric struct {
	ID        string    `json:"id"`
	PlantID   string    `json:"plant_id"`
	Key       string    `json:"key"` // E.g., active_permits_count, carbon_emission_kg
	Value     float64   `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}
