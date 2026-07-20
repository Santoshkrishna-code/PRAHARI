package lotoaudit

import "time"

// Audit tracks scheduled LOTO compliance safety audits in the facility.
type Audit struct {
	ID          string    `json:"id"`
	PlantID     string    `json:"plant_id"`
	AuditorID   string    `json:"auditor_id"`
	AuditedAt   time.Time `json:"audited_at"`
	ResultScore float64   `json:"result_score"`
	Compliance  bool      `json:"compliance"`
	Notes       string    `json:"notes"`
}
