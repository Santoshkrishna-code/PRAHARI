package scheduler

import "time"

// Task defines integration job scheduler metadata.
type Task struct {
	ID        string    `json:"id"`
	JobID     string    `json:"job_id"`
	CronExpr  string    `json:"cron_expr"` // E.g., 0 */2 * * *
	Active    bool      `json:"active"`
	LastRunAt *time.Time `json:"last_run_at,omitempty"`
	NextRunAt time.Time `json:"next_run_at"`
}
