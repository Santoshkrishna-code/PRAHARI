package schedule

import "time"

// ReportSchedule defines recurring triggers to send executive PDF reports.
type ReportSchedule struct {
	ID         string    `json:"id"`
	PlantID    string    `json:"plant_id"`
	ReportType string    `json:"report_type"` // MONTHLY, QUARTERLY
	Recipients []string  `json:"recipients"`
	CronExpr   string    `json:"cron_expr"`
	NextRunAt  time.Time `json:"next_run_at"`
}
