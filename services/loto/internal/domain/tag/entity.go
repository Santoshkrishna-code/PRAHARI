package tag

import "time"

// Tag represents a physical danger/warning tag in the plant registry.
type Tag struct {
	ID          string     `json:"id"`
	TagNumber   string     `json:"tag_number"`
	Status      string     `json:"status"` // AVAILABLE, APPLIED, DAMAGED
	Details     string     `json:"details,omitempty"`
	AssignedTo  string     `json:"assigned_to,omitempty"`
	AppliedAt   *time.Time `json:"applied_at,omitempty"`
}
