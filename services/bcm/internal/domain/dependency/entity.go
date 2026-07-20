package dependency

import "time"

// Mapping tracks upstream / downstream process, asset, and application dependencies.
type Mapping struct {
	ID             string    `json:"id"`
	ProcessID      string    `json:"process_id"`
	DependsOnType  string    `json:"depends_on_type"` // ASSET, APPLICATION, POWER_UTILITY, WATER_UTILITY, SUPPLIER
	DependsOnID    string    `json:"depends_on_id"`
	Criticality    string    `json:"criticality"`     // CRITICAL, HIGH, MEDIUM
	CreatedAt      time.Time `json:"created_at"`
}
