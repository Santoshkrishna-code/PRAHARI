package responseplan

import "time"

// Plan represents an Emergency Response Plan (ERP) for specific plant scenarios.
type Plan struct {
	ID          string    `json:"id"`
	PlantID     string    `json:"plant_id"`
	PlanNumber  string    `json:"plan_number"`
	Title       string    `json:"title"`
	Category    string    `json:"category"`
	Procedures  string    `json:"procedures"` // Markdown / JSON step-by-step SOPs
	Version     string    `json:"version"`
	ApprovedBy  string    `json:"approved_by"`
	ApprovedAt  time.Time `json:"approved_at"`
	CreatedAt   time.Time `json:"created_at"`
}
