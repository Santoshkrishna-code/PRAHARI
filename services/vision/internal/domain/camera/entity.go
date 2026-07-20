package camera

import "time"

// Camera represents an IP camera registry configuration.
type Camera struct {
	ID        string    `json:"id"`
	PlantID   string    `json:"plant_id"`
	Name      string    `json:"name"` // E.g., Reactor Area Camera 2
	IPAddress string    `json:"ip_address"`
	Status    string    `json:"status"` // ONLINE, OFFLINE, UNREACHABLE
	CreatedAt time.Time `json:"created_at"`
}
