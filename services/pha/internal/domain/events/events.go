package events

import "time"

const (
	EventPHACreated             = "pha.created"
	EventHAZOPCompleted         = "hazop.completed"
	EventLOPACompleted          = "lopa.completed"
	EventRecommendationCreated = "recommendation.created"
	EventPHAApproved            = "pha.approved"
	EventPHARevalidated         = "pha.revalidated"
	EventPHAClosed              = "pha.closed"
)

// BaseEvent holds common event payload parameters.
type BaseEvent struct {
	EventID   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	PlantID   string    `json:"plant_id"`
	Timestamp time.Time `json:"timestamp"`
	Payload   any       `json:"payload"`
}
