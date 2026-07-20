package building

import "time"

// Building represents a structural asset within a physical site/plant.
type Building struct {
	ID        string    `json:"id"`
	PlantID   string    `json:"plant_id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Floors    int       `json:"floors"`
	CreatedAt time.Time `json:"created_at"`
}
