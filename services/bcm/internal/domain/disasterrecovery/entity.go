package disasterrecovery

import "time"

// Plan represents an IT/OT Disaster Recovery (DR) execution playbook per NIST SP 800-34.
type Plan struct {
	ID             string    `json:"id"`
	PlanID         string    `json:"plan_id"`
	SystemName     string    `json:"system_name"`
	DRSiteLocation string    `json:"dr_site_location"`
	FailoverType   string    `json:"failover_type"` // AUTOMATED_ACTIVE_PASSIVE, MANUAL_RESTORE
	TargetRTO      float64   `json:"target_rto_hrs"`
	TargetRPO      float64   `json:"target_rpo_hrs"`
	CreatedAt      time.Time `json:"created_at"`
}
