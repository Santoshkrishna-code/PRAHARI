package alert

import "time"

// Rule defines metric thresholds that trigger alerts when exceeded.
type Rule struct {
	ID        string    `json:"id"`
	MetricKey string    `json:"metric_key"`
	Threshold float64   `json:"threshold"`
	Operator  string    `json:"operator"` // GREATER_THAN, LESS_THAN, EQUALS
	Active    bool      `json:"active"`
	UpdatedAt time.Time `json:"updated_at"`
}
