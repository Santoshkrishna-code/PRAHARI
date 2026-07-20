package timeline

import (
	"encoding/json"
	"time"
)

// EventType classifies the type of timeline entry.
type EventType string

const (
	EventCreated              EventType = "CREATED"
	EventSubmitted            EventType = "SUBMITTED"
	EventStatusChanged        EventType = "STATUS_CHANGED"
	EventAssigned             EventType = "ASSIGNED"
	EventCommentAdded         EventType = "COMMENT_ADDED"
	EventAttachmentUploaded   EventType = "ATTACHMENT_UPLOADED"
	EventEvidenceCollected    EventType = "EVIDENCE_COLLECTED"
	EventInvestigationStarted EventType = "INVESTIGATION_STARTED"
	EventRootCauseIdentified  EventType = "ROOT_CAUSE_IDENTIFIED"
	EventCAPACreated          EventType = "CAPA_CREATED"
	EventCAPACompleted        EventType = "CAPA_COMPLETED"
	EventResolved             EventType = "RESOLVED"
	EventClosed               EventType = "CLOSED"
	EventRejected             EventType = "REJECTED"
	EventEscalated            EventType = "ESCALATED"
)

// Event represents an immutable timeline entry recording a single action
// performed on an incident. Timeline events are append-only and never modified.
type Event struct {
	ID          string          `json:"id" db:"id"`
	IncidentID  string          `json:"incident_id" db:"incident_id"`
	EventType   EventType       `json:"event_type" db:"event_type"`
	ActorID     string          `json:"actor_id" db:"actor_id"`
	Description string          `json:"description" db:"description"`
	Metadata    json.RawMessage `json:"metadata,omitempty" db:"metadata"`
	OccurredAt  time.Time       `json:"occurred_at" db:"occurred_at"`
}

// NewEvent constructs an immutable timeline event with the current timestamp.
func NewEvent(incidentID string, eventType EventType, actorID, description string, metadata map[string]string) *Event {
	var metaJSON json.RawMessage
	if metadata != nil {
		metaJSON, _ = json.Marshal(metadata)
	}

	return &Event{
		IncidentID:  incidentID,
		EventType:   eventType,
		ActorID:     actorID,
		Description: description,
		Metadata:    metaJSON,
		OccurredAt:  time.Now(),
	}
}
