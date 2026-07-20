package events

import "time"

const (
	EventDocumentCreated   = "document.created"
	EventDocumentReviewed  = "document.reviewed"
	EventDocumentApproved  = "document.approved"
	EventDocumentPublished = "document.published"
	EventDocumentRevised   = "document.revised"
	EventDocumentArchived  = "document.archived"
)

// BaseEvent holds common event payload parameters.
type BaseEvent struct {
	EventID   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	PlantID   string    `json:"plant_id"`
	Timestamp time.Time `json:"timestamp"`
	Payload   any       `json:"payload"`
}
