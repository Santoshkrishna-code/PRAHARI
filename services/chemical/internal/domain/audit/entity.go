package audit

import "time"

// Record represents a full compliance audit on regulatory chemical volumes.
type Record struct {
	ID        string    `json:"id"`
	PlantID   string    `json:"plant_id"`
	AuditorID string    `json:"auditor_id"`
	AuditedAt time.Time `json:"audited_at"`
	Score     float64   `json:"score"`
	Comments  string    `json:"comments,omitempty"`
}
