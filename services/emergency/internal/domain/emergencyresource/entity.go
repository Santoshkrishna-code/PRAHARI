package emergencyresource

import "time"

// Resource represents fire tenders, foam monitors, SCBA sets, ambulances, or emergency generators.
type Resource struct {
	ID           string    `json:"id"`
	PlantID      string    `json:"plant_id"`
	ResourceCode string    `json:"resource_code"`
	ResourceType string    `json:"resource_type"`
	Quantity     int       `json:"quantity"`
	AvailableQty int       `json:"available_qty"`
	LocationCode string    `json:"location_code"`
	Status       string    `json:"status"` // READY, DEPLOYED, MAINTENANCE
	CreatedAt    time.Time `json:"created_at"`
}
