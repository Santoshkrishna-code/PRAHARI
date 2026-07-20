package timeline

import (
	"encoding/json"
	"time"
)

// EventType defines valid state transition milestone tags.
type EventType string

const (
	EventCreated           EventType = "HAZARD_CREATED"
	EventAssessed          EventType = "HAZARD_ASSESSED"
	EventMitigated         EventType = "HAZARD_MITIGATED"
	EventApproved          EventType = "HAZARD_APPROVED"
	EventImplemented       EventType = "HAZARD_IMPLEMENTED"
	EventVerified          EventType = "HAZARD_VERIFIED"
	EventClosed            EventType = "HAZARD_CLOSED"
	EventRejected          EventType = "HAZARD_REJECTED"
	EventDeleted           EventType = "HAZARD_DELETED"
	EventControlMeasureAdd EventType = "HAZARD_CONTROL_MEASURE_ADDED"
)

// Event records a timeline milestone in walkthroughs history tracking.
type Event struct {
	ID          string          `json:"id" db:"id"`
	HazardID    string          `json:"hazard_id" db:"hazard_id"`
	EventType   EventType       `json:"event_type" db:"event_type"`
	ActorID     string          `json:"actor_id" db:"actor_id"`
	Description string          `json:"description" db:"description"`
	Metadata    json.RawMessage `json:"metadata,omitempty" db:"metadata"`
	OccurredAt  time.Time       `json:"occurred_at" db:"occurred_at"`
}

// NewEvent instantiates a milestone Event.
func NewEvent(hazardID string, eventType EventType, actorID, description string, metadata map[string]string) *Event {
	var metaJSON json.RawMessage
	if metadata != nil {
		metaJSON, _ = json.Marshal(metadata)
	}

	return &Event{
		HazardID:    hazardID,
		EventType:   eventType,
		ActorID:     actorID,
		Description: description,
		Metadata:    metaJSON,
		OccurredAt:  time.Now(),
	}
}
