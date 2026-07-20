package recoveryobjective

import "time"

// Target defines specific RTO, RPO, and MBCO (Minimum Business Continuity Objective) targets.
type Target struct {
	ID          string    `json:"id"`
	ProcessID   string    `json:"process_id"`
	TargetRTO   float64   `json:"target_rto_hrs"`
	TargetRPO   float64   `json:"target_rpo_hrs"`
	TargetMBCO  float64   `json:"target_mbco_pct"` // E.g., 50% capacity within 24h
	CreatedAt   time.Time `json:"created_at"`
}
