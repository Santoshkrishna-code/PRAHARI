package host

import "time"

// Host represents the plant employee hosting the visitor.
type Host struct {
	ID           string    `json:"id"`
	PlantID      string    `json:"plant_id"`
	HostName     string    `json:"host_name"`
	Department   string    `json:"department"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	CreatedAt    time.Time `json:"created_at"`
}
