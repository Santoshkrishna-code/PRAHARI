package lock

import "time"

// Lock represents a physical padlock in the plant registry.
type Lock struct {
	ID         string     `json:"id"`
	LockNumber string     `json:"lock_number"`
	Color      string     `json:"color"` // Blue (Electrical), Red (Mechanical), Yellow (Group Lockbox)
	Status     string     `json:"status"` // AVAILABLE, APPLIED, LOST
	AssignedTo string     `json:"assigned_to,omitempty"` // UserID
	AppliedAt  *time.Time `json:"applied_at,omitempty"`
}
