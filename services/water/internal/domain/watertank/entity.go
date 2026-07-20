package watertank

import "time"

// Tank represents a physical water storage tank.
type Tank struct {
	ID             string    `json:"id"`
	PlantID        string    `json:"plant_id"`
	TankName       string    `json:"tank_name"`
	TankType       string    `json:"tank_type"`
	MaxCapacityKL  float64   `json:"max_capacity_kl"`
	CurrentLevelKL float64   `json:"current_level_kl"`
	AssetID        string    `json:"asset_id,omitempty"`
	LocationCode   string    `json:"location_code"`
	IsActive       bool      `json:"is_active"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
