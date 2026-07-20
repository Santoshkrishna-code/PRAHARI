package timeline

import (
	"encoding/json"
	"time"
)

// EventType defines valid state/milestone transitions in permit timeline.
type EventType string

const (
	EventCreated             EventType = "PERMIT_CREATED"
	EventSubmitted           EventType = "PERMIT_SUBMITTED"
	EventRiskAssessed        EventType = "RISK_ASSESSED"
	EventApproved            EventType = "APPROVED"
	EventIssued              EventType = "ISSUED"
	EventActivated           EventType = "ACTIVATED"
	EventSuspended           EventType = "SUSPENDED"
	EventExtended            EventType = "EXTENDED"
	EventCompleted           EventType = "COMPLETED"
	EventClosed              EventType = "CLOSED"
	EventCancelled           EventType = "CANCELLED"
	EventGasTestRecorded     EventType = "GAS_TEST_RECORDED"
	EventIsolationApplied    EventType = "ISOLATION_APPLIED"
	EventIsolationRemoved    EventType = "ISOLATION_REMOVED"
	EventCommentAdded        EventType = "COMMENT_ADDED"
	EventAttachmentUploaded  EventType = "ATTACHMENT_UPLOADED"
)

// Event records a milestone action in the permit lifecycle.
type Event struct {
	ID          string          `json:"id" db:"id"`
	PermitID    string          `json:"permit_id" db:"permit_id"`
	EventType   EventType       `json:"event_type" db:"event_type"`
	ActorID     string          `json:"actor_id" db:"actor_id"`
	Description string          `json:"description" db:"description"`
	Metadata    json.RawMessage `json:"metadata,omitempty" db:"metadata"`
	OccurredAt  time.Time       `json:"occurred_at" db:"occurred_at"`
}

// NewEvent constructs an immutable timeline event.
func NewEvent(permitID string, eventType EventType, actorID, description string, metadata map[string]string) *Event {
	var metaJSON json.RawMessage
	if metadata != nil {
		metaJSON, _ = json.Marshal(metadata)
	}

	return &Event{
		PermitID:    permitID,
		EventType:   eventType,
		ActorID:     actorID,
		Description: description,
		Metadata:    metaJSON,
		OccurredAt:  time.Now(),
	}
}
