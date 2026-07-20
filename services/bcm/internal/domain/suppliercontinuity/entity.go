package suppliercontinuity

import "time"

// Evaluation tracks third-party supplier business continuity readiness and SLA alignment.
type Evaluation struct {
	ID              string    `json:"id"`
	SupplierID      string    `json:"supplier_id"`
	SupplierName    string    `json:"supplier_name"`
	ResilienceScore float64   `json:"resilience_score"` // 0 to 100%
	SLARTOHrs       float64   `json:"sla_rto_hrs"`
	LastAuditedAt   time.Time `json:"last_audited_at"`
	CreatedAt       time.Time `json:"created_at"`
}
