package recoveryteam

import "time"

// Team represents a designated Business Continuity & Disaster Recovery Team.
type Team struct {
	ID        string    `json:"id"`
	PlanID    string    `json:"plan_id"`
	TeamName  string    `json:"team_name"`
	LeaderID  string    `json:"leader_id"`
	Role      string    `json:"role"` // BCM_COMMAND, IT_DR_RECOVERY, LOGISTICS_RECOVERY
	CreatedAt time.Time `json:"created_at"`
}
