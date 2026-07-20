package timeline

import (
	"encoding/json"
	"time"
)

// EventType defines valid state transition milestone tags.
type EventType string

const (
	EventCreated      EventType = "MAINTENANCE_CREATED"
	EventApproved     EventType = "MAINTENANCE_APPROVED"
	EventScheduled    EventType = "MAINTENANCE_SCHEDULED"
	EventAssigned     EventType = "MAINTENANCE_ASSIGNED"
	EventStarted      EventType = "MAINTENANCE_STARTED"
	EventCompleted    EventType = "MAINTENANCE_COMPLETED"
	EventVerified     EventType = "MAINTENANCE_VERIFIED"
	EventClosed       EventType = "MAINTENANCE_CLOSED"
	EventCancelled    EventType = "MAINTENANCE_CANCELLED"
	EventPartUsed     EventType = "MAINTENANCE_PART_USED"
	EventLaborLogged  EventType = "MAINTENANCE_LABOR_LOGGED"
	EventDowntimeRecorded EventType = "MAINTENANCE_DOWNTIME_RECORDED"
)

// Event records a timeline milestone in walkthroughs history tracking.
type Event struct {
	ID            string          `json:"id" db:"id"`
	MaintenanceID string          `json:"maintenance_id" db:"maintenance_id"`
	EventType     EventType       `json:"event_type" db:"event_type"`
	ActorID       string          `json:"actor_id" db:"actor_id"`
	Description   string          `json:"description" db:"description"`
	Metadata      json.RawMessage `json:"metadata,omitempty" db:"metadata"`
	OccurredAt    time.Time       `json:"occurred_at" db:"occurred_at"`
}

// NewEvent instantiates a milestone Event.
func NewEvent(maintenanceID string, eventType EventType, actorID, description string, metadata map[string]string) *Event {
	var metaJSON json.RawMessage
	if metadata != nil {
		metaJSON, _ = json.Marshal(metadata)
	}

	return &Event{
		MaintenanceID: maintenanceID,
		EventType:     eventType,
		ActorID:       actorID,
		Description:   description,
		Metadata:      metaJSON,
		OccurredAt:    time.Now(),
	}
}
