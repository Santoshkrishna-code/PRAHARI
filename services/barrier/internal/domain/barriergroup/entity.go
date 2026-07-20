package barriergroup

import "time"

// Group represents a logical cluster of barriers protecting a specific process element.
type Group struct {
	ID          string    `json:"id"`
	PlantID     string    `json:"plant_id"`
	GroupName   string    `json:"group_name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
