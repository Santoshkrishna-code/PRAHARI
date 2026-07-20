package drill

import "time"

// Drill represents an emergency drill execution (Mock Evacuation, Fire Drill, Chemical Spill Drill).
type Drill struct {
	ID          string    `json:"id"`
	PlantID     string    `json:"plant_id"`
	Title       string    `json:"title"`
	DrillType   string    `json:"drill_type"` // FIRE, SPILL, EVACUATION
	ScheduledAt time.Time `json:"scheduled_at"`
	ExecutedAt  *time.Time `json:"executed_at,omitempty"`
	DurationMin float64   `json:"duration_min"`
	Passed      bool      `json:"passed"`
	EvaluatorID string    `json:"evaluator_id"`
	Status      string    `json:"status"` // SCHEDULED, IN_PROGRESS, COMPLETED, CANCELLED
	CreatedAt   time.Time `json:"created_at"`
}
