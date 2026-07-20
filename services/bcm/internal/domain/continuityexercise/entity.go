package continuityexercise

import "time"

// Exercise represents a business continuity testing exercise.
type Exercise struct {
	ID          string    `json:"id"`
	PlanID      string    `json:"plan_id"`
	Title       string    `json:"title"`
	Type        string    `json:"type"` // TABLETOP, SIMULATION, FULL_FAILOVER
	ScheduledAt time.Time `json:"scheduled_at"`
	ExecutedAt  *time.Time `json:"executed_at,omitempty"`
	Passed      bool      `json:"passed"`
	RTOAchieved float64   `json:"rto_achieved_hrs"`
	Status      string    `json:"status"` // SCHEDULED, IN_PROGRESS, COMPLETED, CANCELLED
	CreatedAt   time.Time `json:"created_at"`
}
