package timeline

import (
	"encoding/json"
	"time"
)

// EventType defines valid state transition milestone tags.
type EventType string

const (
	EventCreated      EventType = "INSPECTION_CREATED"
	EventScheduled    EventType = "INSPECTION_SCHEDULED"
	EventAssigned     EventType = "INSPECTION_ASSIGNED"
	EventStarted      EventType = "INSPECTION_STARTED"
	EventCompleted    EventType = "INSPECTION_COMPLETED"
	EventApproved     EventType = "INSPECTION_APPROVED"
	EventClosed       EventType = "INSPECTION_CLOSED"
	EventCancelled    EventType = "INSPECTION_CANCELLED"
	EventFindingCreated EventType = "FINDING_CREATED"
	EventNCRCreated    EventType = "NCR_CREATED"
	EventCAPACreated   EventType = "CAPA_CREATED"
)

// Event records a timeline milestone in walkthroughs history tracking.
type Event struct {
	ID           string          `json:"id" db:"id"`
	InspectionID string          `json:"inspection_id" db:"inspection_id"`
	EventType    EventType       `json:"event_type" db:"event_type"`
	ActorID      string          `json:"actor_id" db:"actor_id"`
	Description  string          `json:"description" db:"description"`
	Metadata     json.RawMessage `json:"metadata,omitempty" db:"metadata"`
	OccurredAt   time.Time       `json:"occurred_at" db:"occurred_at"`
}

// NewEvent instantiates a milestone Event.
func NewEvent(inspectionID string, eventType EventType, actorID, description string, metadata map[string]string) *Event {
	var metaJSON json.RawMessage
	if metadata != nil {
		metaJSON, _ = json.Marshal(metadata)
	}

	return &Event{
		InspectionID: inspectionID,
		EventType:    eventType,
		ActorID:      actorID,
		Description:  description,
		Metadata:     metaJSON,
		OccurredAt:   time.Now(),
	}
}
