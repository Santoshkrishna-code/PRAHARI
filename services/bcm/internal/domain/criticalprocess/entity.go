package criticalprocess

import "time"

// Process represents a mission-critical business or manufacturing process.
type Process struct {
	ID          string    `json:"id"`
	PlantID     string    `json:"plant_id"`
	ProcessCode string    `json:"process_code"`
	ProcessName string    `json:"process_name"`
	OwnerID     string    `json:"owner_id"`
	PriorityTier string   `json:"priority_tier"` // TIER_1_CRITICAL, TIER_2_IMPORTANT, TIER_3_NORMAL
	CreatedAt   time.Time `json:"created_at"`
}
