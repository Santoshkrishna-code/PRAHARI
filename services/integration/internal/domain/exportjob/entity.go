package exportjob

import "time"

// Job represents a batch file export task.
type Job struct {
	ID        string    `json:"id"`
	Format    string    `json:"format"` // CSV, PDF, EXCEL
	Status    string    `json:"status"` // PENDING, PROCESSING, COMPLETED, FAILED
	FileURL   string    `json:"file_url,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
