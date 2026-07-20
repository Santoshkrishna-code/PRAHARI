package timeline

import (
	"encoding/json"
	"time"
)

// EventType defines valid state transition milestone tags.
type EventType string

const (
	EventCreated               EventType = "ASSET_CREATED"
	EventCommissioned         EventType = "ASSET_COMMISSIONED"
	EventOperational          EventType = "ASSET_OPERATIONAL"
	EventMaintenanceStarted   EventType = "ASSET_MAINTENANCE_STARTED"
	EventMaintenanceCompleted EventType = "ASSET_MAINTENANCE_COMPLETED"
	EventRetired               EventType = "ASSET_RETIRED"
	EventDisposed              EventType = "ASSET_DISPOSED"
	EventDeleted               EventType = "ASSET_DELETED"
	EventLocationChanged       EventType = "ASSET_LOCATION_CHANGED"
	EventOwnerChanged          EventType = "ASSET_OWNER_CHANGED"
	EventCriticalityChanged    EventType = "ASSET_CRITICALITY_CHANGED"
)

// Event records a timeline milestone in walkthroughs history tracking.
type Event struct {
	ID          string          `json:"id" db:"id"`
	AssetID     string          `json:"asset_id" db:"asset_id"`
	EventType   EventType       `json:"event_type" db:"event_type"`
	ActorID     string          `json:"actor_id" db:"actor_id"`
	Description string          `json:"description" db:"description"`
	Metadata    json.RawMessage `json:"metadata,omitempty" db:"metadata"`
	OccurredAt  time.Time       `json:"occurred_at" db:"occurred_at"`
}

// NewEvent instantiates a milestone Event.
func NewEvent(assetID string, eventType EventType, actorID, description string, metadata map[string]string) *Event {
	var metaJSON json.RawMessage
	if metadata != nil {
		metaJSON, _ = json.Marshal(metadata)
	}

	return &Event{
		AssetID:     assetID,
		EventType:   eventType,
		ActorID:     actorID,
		Description: description,
		Metadata:    metaJSON,
		OccurredAt:  time.Now(),
	}
}
