package state

import "time"

// LiveState represents the current real-time state of an equipment node.
type LiveState struct {
	ID          string    `json:"id"`
	TwinID      string    `json:"twin_id"`
	EquipmentID string    `json:"equipment_id"`
	Value       float64   `json:"value"`
	Quality     string    `json:"quality"` // GOOD, BAD, UNCERTAIN
	Timestamp   time.Time `json:"timestamp"`
}
