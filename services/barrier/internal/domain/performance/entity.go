package performance

import "time"

// Metric represents barrier uptime, reliability, and proof test compliance KPIs.
type Metric struct {
	BarrierID         string    `json:"barrier_id"`
	AvailabilityPct   float64   `json:"availability_pct"`
	ProofTestCompliance float64 `json:"proof_test_compliance_pct"`
	BypassHours       float64   `json:"bypass_hours"`
	ImpairmentHours   float64   `json:"impairment_hours"`
	UpdatedAt         time.Time `json:"updated_at"`
}
