package failovertest

import "time"

// Test tracks technical DR failover and failback testing results.
type Test struct {
	ID          string    `json:"id"`
	DRPlanID    string    `json:"dr_plan_id"`
	TestName    string    `json:"test_name"`
	Passed      bool      `json:"passed"`
	RTOAchieved float64   `json:"rto_achieved_hrs"`
	RPOAchieved float64   `json:"rpo_achieved_hrs"`
	ExecutedAt  time.Time `json:"executed_at"`
}
