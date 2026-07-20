package recommendation

import "time"

// Recommendation represents AI-derived corrective actions suggestions.
type Recommendation struct {
	ID        string    `json:"id"`
	PlantID   string    `json:"plant_id"`
	Type      string    `json:"type"` // E.g., preventive_maintenance, safety_training
	SourceID  string    `json:"source_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
