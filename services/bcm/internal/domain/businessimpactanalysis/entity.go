package businessimpactanalysis

import "time"

// Analysis represents a Business Impact Analysis (BIA) evaluating operational, financial, and regulatory impacts.
type Analysis struct {
	ID                 string    `json:"id"`
	PlanID             string    `json:"plan_id"`
	ProcessID          string    `json:"process_id"`
	FinancialLossPerDay float64  `json:"financial_loss_per_day"`
	OperationalImpact  string    `json:"operational_impact"` // SEVERE, MODERATE, MINOR
	RegulatoryImpact   string    `json:"regulatory_impact"`  // NON_COMPLIANCE_FINE, STATUTORY_SHUTDOWN
	MaximumTolerableDowntimeHrs float64 `json:"maximum_tolerable_downtime_hrs"` // MTD / MTPD
	RTOHrs             float64   `json:"rto_hrs"`           // Recovery Time Objective
	RPOHrs             float64   `json:"rpo_hrs"`           // Recovery Point Objective
	EvaluatedAt        time.Time `json:"evaluated_at"`
	CreatedAt          time.Time `json:"created_at"`
}
