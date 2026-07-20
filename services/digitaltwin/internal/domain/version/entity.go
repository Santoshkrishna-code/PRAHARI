package version

import "time"

// Snapshot represents a versioned snapshot of a digital twin state.
type Snapshot struct {
	ID        string    `json:"id"`
	TwinID    string    `json:"twin_id"`
	Version   int       `json:"version"`
	Label     string    `json:"label"` // E.g., v2.1-pre-turnaround
	StateData string    `json:"state_data"` // JSON serialized twin state
	CreatedAt time.Time `json:"created_at"`
}
