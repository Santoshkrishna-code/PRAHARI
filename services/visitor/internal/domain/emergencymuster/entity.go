package emergencymuster

import "time"

// Record tracks active visitors status during emergency evacuation/muster rolls.
type Record struct {
	ID             string     `json:"id"`
	MusterID       string     `json:"muster_id"`
	VisitorID      string     `json:"visitor_id"`
	AssemblyPoint  string     `json:"assembly_point,omitempty"`
	AccountedFor   bool       `json:"accounted_for"`
	AccountedAt    *time.Time `json:"accounted_at,omitempty"`
	WardenID       string     `json:"warden_id,omitempty"`
}
