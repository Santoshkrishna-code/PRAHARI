package continuitystrategy

import "time"

// Strategy defines technical / operational mitigation strategies to maintain critical processes.
type Strategy struct {
	ID           string    `json:"id"`
	PlanID       string    `json:"plan_id"`
	StrategyType string    `json:"strategy_type"` // REPLICA_SITE, WORK_FROM_HOME, ALTERNATE_SUPPLIER, MANUAL_WORKAROUND
	Description  string    `json:"description"`
	CostEstimate float64   `json:"cost_estimate"`
	Approved     bool      `json:"approved"`
	CreatedAt    time.Time `json:"created_at"`
}
