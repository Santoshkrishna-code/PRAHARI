package timeline

import "time"

// Item represents a BCM timeline milestone.
type Item struct {
	ID          string    `json:"id"`
	RecordID    string    `json:"record_id"`
	EventType   string    `json:"event_type"`
	EventDate   time.Time `json:"event_date"`
	ActorID     string    `json:"actor_id"`
	Description string    `json:"description"`
	Metadata    string    `json:"metadata,omitempty"`
}
