package recovery

import "time"

// Plan represents post-emergency site recovery and restoration tracking.
type Plan struct {
	ID             string     `json:"id"`
	EmergencyID    string     `json:"emergency_id"`
	Title          string     `json:"title"`
	DamageSummary  string     `json:"damage_summary"`
	EstimatedCost  float64    `json:"estimated_cost"`
	Status         string     `json:"status"` // IN_PROGRESS, COMPLETED
	TargetComplete time.Time  `json:"target_complete"`
	CompletedAt    *time.Time `json:"completed_at,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
}
