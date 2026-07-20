package ppemaintenance

import "time"

// Record tracks technical repair, recalibration, or deep sanitization of high-grade PPE items.
type Record struct {
	ID            string     `json:"id"`
	ItemID        string     `json:"item_id"`
	MaintenanceBy string     `json:"maintenance_by"`
	MaintenanceAt time.Time  `json:"maintenance_at"`
	Cost          float64    `json:"cost"`
	ActionsTaken  string     `json:"actions_taken"`
	CompletedAt   *time.Time `json:"completed_at,omitempty"`
}
