package escalation

import "time"

// Rule triggers alerts and notifies higher management levels if action items remain overdue beyond specific thresholds.
type Rule struct {
	ID             string    `json:"id"`
	PlantID        string    `json:"plant_id"`
	OverdueDays    int       `json:"overdue_days"`
	EscalateToRole string    `json:"escalate_to_role"` // E.g., Plant Manager, VP Safety
	NotifyEmail    string    `json:"notify_email"`
	Active         bool      `json:"active"`
	CreatedAt      time.Time `json:"created_at"`
}
