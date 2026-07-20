package importjob

import "time"

// Job represents a batch file import task (CSV, JSON).
type Job struct {
	ID         string    `json:"id"`
	FileName   string    `json:"file_name"`
	TargetArea string    `json:"target_area"` // ASSETS, USERS, CHEMICALS
	Status     string    `json:"status"`      // PENDING, PROCESSING, COMPLETED, FAILED
	Progress   float64   `json:"progress"`
	CreatedAt  time.Time `json:"created_at"`
}
