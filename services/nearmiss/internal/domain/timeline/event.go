package timeline

import (
	"encoding/json"
	"time"
)

// EventType defines valid state transition milestone tags.
type EventType string

const (
	EventCreated      EventType = "NEARMISS_CREATED"
	EventClassified   EventType = "NEARMISS_CLASSIFIED"
	EventInvestigated EventType = "NEARMISS_INVESTIGATED"
	EventCorrectiveAdd EventType = "NEARMISS_CORRECTIVE_ADDED"
	EventVerified     EventType = "NEARMISS_VERIFIED"
	EventClosed       EventType = "NEARMISS_CLOSED"
	EventEscalated    EventType = "NEARMISS_ESCALATED"
	EventDeleted      EventType = "NEARMISS_DELETED"
)

// Event records a timeline milestone in walkthroughs history tracking.
type Event struct {
	ID         string          `json:"id" db:"id"`
	NearMissID string          `json:"near_miss_id" db:"near_miss_id"`
	EventType  EventType       `json:"event_type" db:"event_type"`
	ActorID    string          `json:"actor_id" db:"actor_id"`
	Description string          `json:"description" db:"description"`
	Metadata   json.RawMessage `json:"metadata,omitempty" db:"metadata"`
	OccurredAt time.Time       `json:"occurred_at" db:"occurred_at"`
}

// NewEvent instantiates a milestone Event.
func NewEvent(nearmissID string, eventType EventType, actorID, description string, metadata map[string]string) *Event {
	var metaJSON json.RawMessage
	if metadata != nil {
		metaJSON, _ = json.Marshal(metadata)
	}

	return &Event{
		NearMissID: nearmissID,
		EventType:  eventType,
		ActorID:    actorID,
		Description: description,
		Metadata:   metaJSON,
		OccurredAt: time.Now(),
	}
}
