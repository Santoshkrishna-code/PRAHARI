package emergencyteam

import "time"

// Team represents an Emergency Response Team (ERT, Fire Brigade, Hazmat Team, Rescue Team).
type Team struct {
	ID        string    `json:"id"`
	PlantID   string    `json:"plant_id"`
	TeamName  string    `json:"team_name"`
	TeamType  string    `json:"team_type"` // FIRE, HAZMAT, RESCUE, MEDICAL
	LeaderID  string    `json:"leader_id"`
	CreatedAt time.Time `json:"created_at"`
}
