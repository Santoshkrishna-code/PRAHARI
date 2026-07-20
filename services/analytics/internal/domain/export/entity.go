package export

import "time"

// Job tracks export operations.
type Job struct {
	ID        string    `json:"id"`
	Format    string    `json:"format"` // CSV, PDF, EXCEL
	Status    string    `json:"status"` // PENDING, PROCESSING, COMPLETED, FAILED
	FileURL   string    `json:"file_url,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
