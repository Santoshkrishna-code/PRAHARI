package timeline

import (
	"encoding/json"
	"time"
)

// EventType defines valid state transition milestone tags.
type EventType string

const (
	EventCreated      EventType = "OBSERVATION_CREATED"
	EventReviews      EventType = "OBSERVATION_REVIEWED"
	EventCoached      EventType = "OBSERVATION_COACHED"
	EventRecognized   EventType = "OBSERVATION_RECOGNIZED"
	EventFollowupDone EventType = "OBSERVATION_FOLLOWUP_DONE"
	EventVerified     EventType = "OBSERVATION_VERIFIED"
	EventClosed       EventType = "OBSERVATION_CLOSED"
	EventDeleted      EventType = "OBSERVATION_DELETED"
)

// Event records a timeline milestone in walkthroughs history tracking.
type Event struct {
	ID            string          `json:"id" db:"id"`
	ObservationID string          `json:"observation_id" db:"observation_id"`
	EventType     EventType       `json:"event_type" db:"event_type"`
	ActorID       string          `json:"actor_id" db:"actor_id"`
	Description   string          `json:"description" db:"description"`
	Metadata      json.RawMessage `json:"metadata,omitempty" db:"metadata"`
	OccurredAt    time.Time       `json:"occurred_at" db:"occurred_at"`
}

// NewEvent instantiates a milestone Event.
func NewEvent(observationID string, eventType EventType, actorID, description string, metadata map[string]string) *Event {
	var metaJSON json.RawMessage
	if metadata != nil {
		metaJSON, _ = json.Marshal(metadata)
	}

	return &Event{
		ObservationID: observationID,
		EventType:     eventType,
		ActorID:       actorID,
		Description:   description,
		Metadata:      metaJSON,
		OccurredAt:    time.Now(),
	}
}
}
