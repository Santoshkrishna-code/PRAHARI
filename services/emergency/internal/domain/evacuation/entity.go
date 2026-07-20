package evacuation

import "time"

// Record tracks plant evacuation execution and headcount accountability.
type Record struct {
	ID               string     `json:"id"`
	EmergencyID      string     `json:"emergency_id"`
	ZoneID           string     `json:"zone_id"`
	InitiatedBy      string     `json:"initiated_by"`
	TotalPersonnel   int        `json:"total_personnel"`
	AccountedFor     int        `json:"accounted_for"`
	MissingCount     int        `json:"missing_count"`
	Status           string     `json:"status"` // IN_PROGRESS, COMPLETED
	EvacuationTimeSec float64   `json:"evacuation_time_sec"`
	InitiatedAt      time.Time  `json:"initiated_at"`
	CompletedAt      *time.Time `json:"completed_at,omitempty"`
}
