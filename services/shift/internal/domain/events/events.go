package events

import "time"

const (
	EventShiftStarted      = "shift.started"
	EventHandoverInitiated = "handover.initiated"
	EventHandoverAccepted  = "handover.accepted"
	EventShiftClosed       = "shift.closed"
	EventCrewAssigned      = "crew.assigned"
)

// BaseEvent holds common event payload parameters.
type BaseEvent struct {
	EventID   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	PlantID   string    `json:"plant_id"`
	Timestamp time.Time `json:"timestamp"`
	Payload   any       `json:"payload"`
}
