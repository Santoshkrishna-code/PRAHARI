package storagearea

import "time"

// Area represents a defined physical zone, room, or shelf where chemicals are stored.
type Area struct {
	ID             string    `json:"id"`
	PlantID        string    `json:"plant_id"`
	Name           string    `json:"name"`
	Code           string    `json:"code"`
	VentilationType string   `json:"ventilation_type"` // E.g., NATURAL, FORCED, FUME_HOOD
	MaxCapacityKg  float64   `json:"max_capacity_kg"`
	CurrentLoadKg  float64   `json:"current_load_kg"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
