package report

import "time"

// Report represents a compiled cross-service executive compliance document log.
type Report struct {
	ID         string    `json:"id"`
	PlantID    string    `json:"plant_id"`
	Title      string    `json:"title"`
	ReportType string    `json:"report_type"` // MONTHLY, QUARTERLY, ANNUAL
	FileURL    string    `json:"file_url,omitempty"`
	CreatedBy  string    `json:"created_by"`
	CreatedAt  time.Time `json:"created_at"`
}
