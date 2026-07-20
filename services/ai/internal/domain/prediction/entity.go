package prediction

import "time"

// Forecast represents derived forecast probabilities on compliance trends.
type Forecast struct {
	ID          string    `json:"id"`
	PlantID     string    `json:"plant_id"`
	TargetTopic string    `json:"target_topic"` // E.g., permit_lapses, risk_growth
	Probability float64   `json:"probability"`
	CreatedAt   time.Time `json:"created_at"`
}
