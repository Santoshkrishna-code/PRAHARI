package securitycheckpoint

import "time"

// Checkpoint represents a gate house or access control point on site.
type Checkpoint struct {
	ID             string    `json:"id"`
	PlantID        string    `json:"plant_id"`
	CheckpointName string    `json:"checkpoint_name"`
	Location       string    `json:"location"`
	Active         bool      `json:"active"`
	CreatedAt      time.Time `json:"created_at"`
}
