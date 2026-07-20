package events

import "time"

const (
	EventMOCCreated          = "moc.created"
	EventMOCReviewStarted    = "moc.review.started"
	EventMOCApproved         = "moc.approved"
	EventMOCImplemented      = "moc.implemented"
	EventMOCVerified         = "moc.verified"
	EventMOCRollbackExecuted = "moc.rollback.executed"
	EventMOCCclosed          = "moc.closed"
)

// BaseEvent holds common event payload parameters.
type BaseEvent struct {
	EventID   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	PlantID   string    `json:"plant_id"`
	Timestamp time.Time `json:"timestamp"`
	Payload   any       `json:"payload"`
}
