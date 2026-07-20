package calendar

import "time"

// Calendar represents a shift or operational schedule calendar mapping.
type Calendar struct {
	ID        string    `json:"id"`
	PlantID   string    `json:"plant_id"`
	Name      string    `json:"name"`
	Year      int       `json:"year"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
}
