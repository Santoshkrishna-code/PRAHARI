package synchronization

import "time"

// Record tracks synchronisation instances, counts, and run logs.
type Record struct {
	ID         string     `json:"id"`
	JobID      string     `json:"job_id"`
	StartedAt  time.Time  `json:"started_at"`
	FinishedAt *time.Time `json:"finished_at,omitempty"`
	Status     string     `json:"status"` // RUNNING, SUCCESS, FAILED
	RecordsCount int      `json:"records_count"`
	ErrorMessage string   `json:"error_message,omitempty"`
}
