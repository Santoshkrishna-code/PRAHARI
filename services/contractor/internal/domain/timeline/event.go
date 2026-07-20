package timeline

import (
	"encoding/json"
	"time"
)

// EventType defines valid state transition milestone tags.
type EventType string

const (
	EventCreated      EventType = "CONTRACTOR_CREATED"
	EventApproved     EventType = "CONTRACTOR_APPROVED"
	EventActivated    EventType = "CONTRACTOR_ACTIVATED"
	EventSuspended    EventType = "CONTRACTOR_SUSPENDED"
	EventOffboarded   EventType = "CONTRACTOR_OFFBOARDED"
	EventDeleted      EventType = "CONTRACTOR_DELETED"
	EventTrainingCompleted EventType = "CONTRACTOR_TRAINING_COMPLETED"
	EventMedicalCleared    EventType = "CONTRACTOR_MEDICAL_CLEARED"
	EventSiteAccessGranted EventType = "CONTRACTOR_SITEACCESS_GRANTED"
)

// Event records a timeline milestone in walkthroughs history tracking.
type Event struct {
	ID           string          `json:"id" db:"id"`
	ContractorID string          `json:"contractor_id" db:"contractor_id"`
	EventType    EventType       `json:"event_type" db:"event_type"`
	ActorID      string          `json:"actor_id" db:"actor_id"`
	Description  string          `json:"description" db:"description"`
	Metadata     json.RawMessage `json:"metadata,omitempty" db:"metadata"`
	OccurredAt   time.Time       `json:"occurred_at" db:"occurred_at"`
}

// NewEvent instantiates a milestone Event.
func NewEvent(contractorID string, eventType EventType, actorID, description string, metadata map[string]string) *Event {
	var metaJSON json.RawMessage
	if metadata != nil {
		metaJSON, _ = json.Marshal(metadata)
	}

	return &Event{
		ContractorID: contractorID,
		EventType:    eventType,
		ActorID:      actorID,
		Description:  description,
		Metadata:     metaJSON,
		OccurredAt:   time.Now(),
	}
}
