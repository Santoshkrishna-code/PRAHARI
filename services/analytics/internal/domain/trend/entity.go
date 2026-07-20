package trend

import "time"

// Analysis represents calculated trend line data coordinates.
type Analysis struct {
	ID        string    `json:"id"`
	PlantID   string    `json:"plant_id"`
	MetricKey string    `json:"metric_key"`
	Direction string    `json:"direction"` // UPWARD, DOWNWARD, FLAT
	Slope     float64   `json:"slope"`
	CalculatedAt time.Time `json:"calculated_at"`
}
