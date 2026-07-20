package continuityplan

import "time"

// Plan represents a Business Continuity Plan (BCP) compliant with ISO 22301.
type Plan struct {
	ID          string    `json:"id"`
	PlanNumber  string    `json:"plan_number"`
	PlantID     string    `json:"plant_id"`
	BusinessUnit string   `json:"business_unit"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Scope       string    `json:"scope"`
	Version     string    `json:"version"`
	Status      string    `json:"status"` // Planned, Business Impact Analysis, Strategy Development, Plan Development, Approval, Exercise, Activation, Recovery, Review, Continuous Improvement, Archived
	ApprovedBy  string    `json:"approved_by"`
	ApprovedAt  *time.Time `json:"approved_at,omitempty"`
	NextReviewAt *time.Time `json:"next_review_at,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
