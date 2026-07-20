package timeline

import (
	"time"
)

// Event tracks state changes.
type Event struct {
	ID         string            `json:"id" db:"id"`
	RecordID   string            `json:"record_id" db:"record_id"`
	EventType  string            `json:"event_type" db:"event_type"` // e.g. "CARBON_CALCULATED", "DISCLOSURE_PUBLISHED"
	EventDate  time.Time         `json:"event_date" db:"event_date"`
	ActorID    string            `json:"actor_id" db:"actor_id"`
	Description string           `json:"description" db:"description"`
	Metadata   map[string]string `json:"metadata" db:"metadata"`
}

// NewEvent constructs Event.
func NewEvent(recID, eventType, actorID, desc string, meta map[string]string) *Event {
	return &Event{
		RecordID:    recID,
		EventType:   eventType,
		EventDate:   time.Now(),
		ActorID:     actorID,
		Description: desc,
		Metadata:    meta,
	}
}
