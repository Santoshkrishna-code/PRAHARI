package events

import "time"

const (
	EventLOTOPlanned        = "loto.planned"
	EventIsolationApproved  = "isolation.approved"
	EventLocksApplied       = "locks.applied"
	EventZeroEnergyVerified = "zeroenergy.verified"
	EventSystemRestored     = "system.restored"
	EventLOTOClosed         = "loto.closed"
)

// BaseEvent holds common event payload parameters.
type BaseEvent struct {
	EventID   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	PlantID   string    `json:"plant_id"`
	Timestamp time.Time `json:"timestamp"`
	Payload   any       `json:"payload"`
}
