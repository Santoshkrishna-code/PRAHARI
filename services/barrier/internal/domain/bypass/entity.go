package bypass

import "time"

// Record tracks physical or override bypasses of safety barriers.
type Record struct {
	ID               string     `json:"id"`
	BarrierID        string     `json:"barrier_id"`
	PermitID         string     `json:"permit_id,omitempty"` // Permit-to-work reference
	BypassReason     string     `json:"bypass_reason"`
	ApprovedBy       string     `json:"approved_by"`
	AuthorizedPeriod string     `json:"authorized_period"`
	IsActive         bool       `json:"is_active"`
	BypassedAt       time.Time  `json:"bypassed_at"`
	RestoredAt       *time.Time `json:"restored_at,omitempty"`
}
