package aggregation

import "time"

// Rule defines how raw transactional counts are aggregated (weekly/monthly count sums).
type Rule struct {
	ID          string    `json:"id"`
	MetricKey   string    `json:"metric_key"`
	Interval    string    `json:"interval"` // HOUR, DAY, WEEK, MONTH
	Formula     string    `json:"formula"`  // SUM, AVG, MAX, MIN
	LastRunAt   time.Time `json:"last_run_at"`
}
