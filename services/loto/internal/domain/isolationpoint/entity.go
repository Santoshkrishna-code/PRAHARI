package isolationpoint

import "time"

// Point represents a physical isolation point on an asset/system (e.g. breaker, valve, blind flange).
type Point struct {
	ID             string    `json:"id"`
	EquipmentID    string    `json:"equipment_id"`
	SourceID       string    `json:"source_id"`
	PointName      string    `json:"point_name"` // E.g., Valve V-102, Breaker CB-12
	Location       string    `json:"location"`
	IsolationMethod string   `json:"isolation_method"` // LOCK, TAG, BLIND, DOUBLE_BLOCK_AND_BLEED
	CreatedAt      time.Time `json:"created_at"`
}
