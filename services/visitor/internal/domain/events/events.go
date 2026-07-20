package events

import "time"

const (
	EventVisitorRegistered    = "visitor.registered"
	EventVisitorApproved      = "visitor.approved"
	EventVisitorCheckedIn     = "visitor.checkedin"
	EventVisitorCheckedOut    = "visitor.checkedout"
	EventVisitorBlacklisted   = "visitor.blacklisted"
	EventVisitorMusterComplete = "visitor.muster.completed"
)

// BaseEvent holds common event payload parameters.
type BaseEvent struct {
	EventID   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	PlantID   string    `json:"plant_id"`
	Timestamp time.Time `json:"timestamp"`
	Payload   any       `json:"payload"`
}
