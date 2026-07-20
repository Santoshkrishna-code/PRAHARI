package events

import "time"

const (
	EventBarrierCreated        = "barrier.created"
	EventBarrierIntegrityFailed = "barrier.integrity.failed"
	EventBarrierBypassed       = "barrier.bypassed"
	EventProofTestCompleted    = "prooftest.completed"
	EventBarrierRestored       = "barrier.restored"
	EventBarrierRetired        = "barrier.retired"
)

// BaseEvent holds common event payload parameters.
type BaseEvent struct {
	EventID   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	PlantID   string    `json:"plant_id"`
	Timestamp time.Time `json:"timestamp"`
	Payload   any       `json:"payload"`
}
