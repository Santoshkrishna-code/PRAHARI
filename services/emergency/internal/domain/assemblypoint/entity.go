package assemblypoint

import "time"

// Point represents a safe designated evacuation assembly point.
type Point struct {
	ID             string    `json:"id"`
	PlantID        string    `json:"plant_id"`
	PointCode      string    `json:"point_code"`
	Name           string    `json:"name"`
	Location       string    `json:"location"`
	Capacity       int       `json:"capacity"`
	CurrentCount   int       `json:"current_count"`
	WardenUserID   string    `json:"warden_user_id"`
	CreatedAt      time.Time `json:"created_at"`
}
