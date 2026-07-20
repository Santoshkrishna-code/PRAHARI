package alarm

import "time"

// Visualization represents an alarm mapped to a twin location for display.
type Visualization struct {
	ID          string    `json:"id"`
	TwinID      string    `json:"twin_id"`
	EquipmentID string    `json:"equipment_id"`
	Severity    string    `json:"severity"` // CRITICAL, HIGH, MEDIUM, LOW
	Message     string    `json:"message"`
	Active      bool      `json:"active"`
	TriggeredAt time.Time `json:"triggered_at"`
}
