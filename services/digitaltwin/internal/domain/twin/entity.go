package twin

import "time"

// DigitalTwin represents a registered digital twin instance for a plant or facility.
type DigitalTwin struct {
	ID        string    `json:"id"`
	PlantID   string    `json:"plant_id"`
	Name      string    `json:"name"` // E.g., Refinery Unit 3 Digital Twin
	Status    string    `json:"status"` // DRAFT, ACTIVE, ARCHIVED
	Version   int       `json:"version"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
