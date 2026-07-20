package timeline

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// EventType defines valid state transition milestone tags.
type EventType string

const (
	EventCreated           EventType = "TRAINING_CREATED"
	EventScheduled         EventType = "TRAINING_SCHEDULED"
	EventEnrolled          EventType = "TRAINING_ENROLLED"
	EventStarted           EventType = "TRAINING_STARTED"
	EventCompleted         EventType = "TRAINING_COMPLETED"
	EventCertified         EventType = "TRAINING_CERTIFIED"
	EventExpired           EventType = "TRAINING_EXPIRED"
	EventDeleted           EventType = "TRAINING_DELETED"
)

// Event records a timeline milestone in walkthroughs history tracking.
type Event struct {
	ID          string          `json:"id" db:"id"`
	TrainingID  string          `json:"training_id" db:"training_id"`
	EventType   EventType       `json:"event_type" db:"event_type"`
	ActorID     string          `json:"actor_id" db:"actor_id"`
	Description string          `json:"description" db:"description"`
	Metadata    json.RawMessage `json:"metadata,omitempty" db:"metadata"`
	OccurredAt  time.Time       `json:"occurred_at" db:"occurred_at"`
}

// NewEvent instantiates a milestone Event.
func NewEvent(trainingID string, eventType EventType, actorID, description string, metadata map[string]string) *Event {
	var metaJSON json.RawMessage
	if metadata != nil {
		metaJSON, _ = json.Marshal(metadata)
	}

	return &Event{
		ID:          uuid.New().String(),
		TrainingID:  trainingID,
		EventType:   eventType,
		ActorID:     actorID,
		Description: description,
		Metadata:    metaJSON,
		OccurredAt:  time.Now(),
	}
}

