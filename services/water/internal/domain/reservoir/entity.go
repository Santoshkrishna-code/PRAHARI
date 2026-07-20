package reservoir

import "time"

// Reservoir represents a natural or constructed water impoundment.
type Reservoir struct {
	ID              string    `json:"id"`
	PlantID         string    `json:"plant_id"`
	ReservoirName   string    `json:"reservoir_name"`
	MaxCapacityKL   float64   `json:"max_capacity_kl"`
	CurrentLevelKL  float64   `json:"current_level_kl"`
	MinOperatingKL  float64   `json:"min_operating_kl"`
	LocationCode    string    `json:"location_code"`
	LastInspectedAt time.Time `json:"last_inspected_at"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
