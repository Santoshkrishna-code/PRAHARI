package responder

import "time"

// Person represents a certified responder deployed to an emergency.
type Person struct {
	ID          string    `json:"id"`
	TeamID      string    `json:"team_id"`
	UserID      string    `json:"user_id"`
	Role        string    `json:"role"`
	Status      string    `json:"status"` // STANDBY, DEPLOYED, DISCHARGED
	DeployedAt  time.Time `json:"deployed_at"`
}
