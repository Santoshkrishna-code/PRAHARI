package generatedaction

import "time"

// Action represents an action item generated from a meeting decision that will be
// pushed to the central Action Tracking & CAPA Management Service.
type Action struct {
	ID           string    `json:"id"`
	MeetingID    string    `json:"meeting_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	AssignedTo   string    `json:"assigned_to"`
	DueDate      time.Time `json:"due_date"`
	Priority     string    `json:"priority"` // LOW, MEDIUM, HIGH, CRITICAL
	SyncedToCAPA bool      `json:"synced_to_capa"`
	CAPARefID    string    `json:"capa_ref_id,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}
