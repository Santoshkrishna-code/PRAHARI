package timeline

import (
	"encoding/json"
	"time"
)

// EventType defines valid state transition milestone tags.
type EventType string

const (
	EventCreated           EventType = "COMPLIANCE_CREATED"
	EventReviewed          EventType = "COMPLIANCE_REVIEWED"
	EventEvidenceCollected EventType = "COMPLIANCE_EVIDENCE_COLLECTED"
	EventCompliant         EventType = "COMPLIANCE_COMPLIANT"
	EventNonCompliant      EventType = "COMPLIANCE_NONCOMPLIANT"
	EventExpired           EventType = "COMPLIANCE_EXPIRED"
	EventDeleted           EventType = "COMPLIANCE_DELETED"
)

// Event records a timeline milestone in walkthroughs history tracking.
type Event struct {
	ID           string          `json:"id" db:"id"`
	ComplianceID string          `json:"compliance_id" db:"compliance_id"`
	EventType    EventType       `json:"event_type" db:"event_type"`
	ActorID      string          `json:"actor_id" db:"actor_id"`
	Description  string          `json:"description" db:"description"`
	Metadata     json.RawMessage `json:"metadata,omitempty" db:"metadata"`
	OccurredAt   time.Time       `json:"occurred_at" db:"occurred_at"`
}

// NewEvent instantiates a milestone Event.
func NewEvent(complianceID string, eventType EventType, actorID, description string, metadata map[string]string) *Event {
	var metaJSON json.RawMessage
	if metadata != nil {
		metaJSON, _ = json.Marshal(metadata)
	}

	return &Event{
		ComplianceID: complianceID,
		EventType:    eventType,
		ActorID:      actorID,
		Description:  description,
		Metadata:     metaJSON,
		OccurredAt:   time.Now(),
	}
}
