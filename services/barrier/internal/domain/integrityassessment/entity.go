package integrityassessment

import "time"

// Assessment evaluates holistic barrier health and operational readiness.
type Assessment struct {
	ID           string    `json:"id"`
	BarrierID    string    `json:"barrier_id"`
	EvaluatorID  string    `json:"evaluator_id"`
	HealthScore  float64   `json:"health_score"` // 0 to 100%
	Status       string    `json:"status"`       // HEALTHY, DEGRADED, CRITICAL
	ActionNeeded string    `json:"action_needed,omitempty"`
	AssessedAt   time.Time `json:"assessed_at"`
}
