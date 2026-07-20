package timeline

import (
	"encoding/json"
	"time"
)

// EventType defines valid state transition milestone tags.
type EventType string

const (
	EventCreated    EventType = "RISK_CREATED"
	EventAssessed   EventType = "RISK_ASSESSED"
	EventReviewed   EventType = "RISK_REVIEWED"
	EventApproved   EventType = "RISK_APPROVED"
	EventActivated  EventType = "RISK_ACTIVATED"
	EventReassessed EventType = "RISK_REASSESSED"
	EventClosed     EventType = "RISK_CLOSED"
	EventDeleted    EventType = "RISK_DELETED"
)

// Event records a timeline milestone in walkthroughs history tracking.
type Event struct {
	ID            string          `json:"id" db:"id"`
	RiskID        string          `json:"risk_id" db:"risk_id"`
	EventType     EventType       `json:"event_type" db:"event_type"`
	ActorID       string          `json:"actor_id" db:"actor_id"`
	Description   string          `json:"description" db:"description"`
	Metadata      json.RawMessage `json:"metadata,omitempty" db:"metadata"`
	OccurredAt    time.Time       `json:"occurred_at" db:"occurred_at"`
}

// NewEvent instantiates a milestone Event.
func NewEvent(riskID string, eventType EventType, actorID, description string, metadata map[string]string) *Event {
	var metaJSON json.RawMessage
	if metadata != nil {
		metaJSON, _ = json.Marshal(metadata)
	}

	return &Event{
		RiskID:      riskID,
		EventType:   eventType,
		ActorID:     actorID,
		Description: description,
		Metadata:    metaJSON,
		OccurredAt:  time.Now(),
	}
}
