package facility

import "time"

// Model represents a physical facility (building, area, floor) in the digital twin.
type Model struct {
	ID        string    `json:"id"`
	TwinID    string    `json:"twin_id"`
	Name      string    `json:"name"` // E.g., Reactor Building A
	Type      string    `json:"type"` // BUILDING, AREA, FLOOR, UNIT
	ParentID  string    `json:"parent_id,omitempty"`
	Latitude  float64   `json:"latitude,omitempty"`
	Longitude float64   `json:"longitude,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
