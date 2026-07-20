package events

import "time"

const (
	EventBCMPlanCreated        = "bcm.plan.created"
	EventBIACompleted          = "bia.completed"
	EventContinuityActivated   = "continuity.activated"
	EventRecoveryStarted       = "recovery.started"
	EventContinuityReviewDone  = "continuity.review.completed"
	EventBCMClosed             = "bcm.closed"
)

// BaseEvent holds common event payload parameters.
type BaseEvent struct {
	EventID   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	PlantID   string    `json:"plant_id"`
	Timestamp time.Time `json:"timestamp"`
	Payload   any       `json:"payload"`
}
