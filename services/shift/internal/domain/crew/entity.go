package crew

import "time"

// Crew represents a designated team of operators, technicians, and supervisors.
type Crew struct {
	ID          string    `json:"id"`
	PlantID     string    `json:"plant_id"`
	CrewName    string    `json:"crew_name"` // Crew A, Crew B, Crew C etc.
	LeadID      string    `json:"lead_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
