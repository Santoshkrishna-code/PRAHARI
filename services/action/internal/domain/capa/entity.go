package capa

import "time"

// Record tracks a CAPA (Corrective and Preventive Action) governance envelope enclosing multiple actions.
type Record struct {
	ID             string    `json:"id"`
	PlantID        string    `json:"plant_id"`
	SourceType     string    `json:"source_type"` // INCIDENT, AUDIT, COMPLIANCE
	SourceID       string    `json:"source_id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Severity       string    `json:"severity"` // CRITICAL, MAJOR, MINOR
	Status         string    `json:"status"`   // OPEN, VERIFICATION, CLOSED
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
