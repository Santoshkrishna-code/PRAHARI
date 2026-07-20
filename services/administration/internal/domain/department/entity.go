package department

import "time"

// Department represents a functional division within a plant.
type Department struct {
	ID        string    `json:"id"`
	PlantID   string    `json:"plant_id"`
	Name      string    `json:"name"`
	ManagerID string    `json:"manager_id"`
	CreatedAt time.Time `json:"created_at"`
}
