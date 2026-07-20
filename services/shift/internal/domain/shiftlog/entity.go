package shiftlog

import "time"

// Log represents a generic entry in the digital operator logbook.
type Log struct {
	ID          string    `json:"id"`
	ShiftID     string    `json:"shift_id"`
	LoggedByID  string    `json:"logged_by_id"`
	Category    string    `json:"category"` // PROCESS, MAINTENANCE, SAFETY, ENVIRO, EVENTS
	LogEntry    string    `json:"log_entry"`
	Timestamp   time.Time `json:"timestamp"`
	IsCritical  bool      `json:"is_critical"`
	AssetID     string    `json:"asset_id,omitempty"`
}
