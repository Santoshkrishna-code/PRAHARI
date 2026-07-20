package timeline

import (
	"encoding/json"
	"time"
)

// EventType defines valid state transition milestone tags.
type EventType string

const (
	EventCreated           EventType = "AUDIT_CREATED"
	EventScheduled         EventType = "AUDIT_SCHEDULED"
	EventStarted           EventType = "AUDIT_STARTED"
	EventEvidenceCollected EventType = "AUDIT_EVIDENCE_COLLECTED"
	EventReviewed          EventType = "AUDIT_REVIEWED"
	EventApproved          EventType = "AUDIT_APPROVED"
	EventClosed            EventType = "AUDIT_CLOSED"
	EventDeleted           EventType = "AUDIT_DELETED"
)

// Event records a timeline milestone in walkthroughs history tracking.
type Event struct {
	ID          string          `json:"id" db:"id"`
	AuditID     string          `json:"audit_id" db:"audit_id"`
	EventType   EventType       `json:"event_type" db:"event_type"`
	ActorID     string          `json:"actor_id" db:"actor_id"`
	Description string          `json:"description" db:"description"`
	Metadata    json.RawMessage `json:"metadata,omitempty" db:"metadata"`
	OccurredAt  time.Time       `json:"occurred_at" db:"occurred_at"`
}

// NewEvent instantiates a milestone Event.
func NewEvent(auditID string, eventType EventType, actorID, description string, metadata map[string]string) *Event {
	var metaJSON json.RawMessage
	if metadata != nil {
		metaJSON, _ = json.Marshal(metadata)
	}

	return &Event{
		AuditID:     auditID,
		EventType:   eventType,
		ActorID:     actorID,
		Description: description,
		Metadata:    metaJSON,
		OccurredAt:  time.Now(),
	}
}
}
