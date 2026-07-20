package deviation

import "time"

// Deviation represents a process parameter deviation (e.g., HIGH PRESSURE, NO FLOW).
type Deviation struct {
	ID           string    `json:"id"`
	NodeID       string    `json:"node_id"`
	GuideWord    string    `json:"guide_word"` // HIGH, LOW, NO, REVERSE, AS WELL AS, PART OF, OTHER THAN
	Parameter    string    `json:"parameter"`  // FLOW, PRESSURE, TEMPERATURE, LEVEL, COMPOSITION, REACTION
	DeviationName string   `json:"deviation_name"`
	CreatedAt    time.Time `json:"created_at"`
}
