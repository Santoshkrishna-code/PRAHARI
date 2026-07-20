package productionline

import "time"

// ProductionLine represents a manufacturing assembly or packaging line structure.
type ProductionLine struct {
	ID        string    `json:"id"`
	PlantID   string    `json:"plant_id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
}
