package inspection

import "time"

// Record tracks physical / mechanical integrity inspections of barriers.
type Record struct {
	ID           string    `json:"id"`
	BarrierID    string    `json:"barrier_id"`
	InspectorID  string    `json:"inspector_id"`
	Passes       bool      `json:"passes"`
	DefectsFound string    `json:"defects_found,omitempty"`
	InspectedAt  time.Time `json:"inspected_at"`
}
