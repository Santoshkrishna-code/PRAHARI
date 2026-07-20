package copilot

import "time"

// Copilot represents a configured AI assistant persona metadata.
type Copilot struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"` // E.g., SDS Assistant, Incident RCA Copilot
	Role      string    `json:"role"`
	Model     string    `json:"model"`
	CreatedAt time.Time `json:"created_at"`
}
