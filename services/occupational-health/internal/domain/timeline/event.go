package timeline

import (
	"time"
)

// Event tracks state changes.
type Event struct {
	ID              string            `json:"id" db:"id"`
	HealthProfileID string            `json:"health_profile_id" db:"health_profile_id"`
	EventType       string            `json:"event_type" db:"event_type"` // e.g. "EXAM_COMPLETED", "CLEARANCE_GRANTED"
	EventDate       time.Time         `json:"event_date" db:"event_date"`
	ActorID         string            `json:"actor_id" db:"actor_id"`
	Description     string            `json:"description" db:"description"`
	Metadata        map[string]string `json:"metadata" db:"metadata"`
}

// NewEvent constructs an Event.
func NewEvent(profileID, eventType, actorID, desc string, meta map[string]string) *Event {
	return &Event{
		HealthProfileID: profileID,
		EventType:       eventType,
		EventDate:       time.Now(),
		ActorID:         actorID,
		Description:     desc,
		Metadata:        meta,
	}
}
