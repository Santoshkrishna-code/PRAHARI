package events

import (
	"time"
)

// Envelope wraps all event messages dispatched across services to standardize attributes.
type Envelope struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`      // Type of event (e.g. IncidentCreated)
	Source    string    `json:"source"`    // Source service name (e.g. auth-service)
	Timestamp time.Time `json:"timestamp"` // Log generation time
	Version   string    `json:"version"`   // Schema version (e.g. 1.0.0)
	Payload   []byte    `json:"payload"`   // Raw serialized payload data
}
