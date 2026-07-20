package events

import "time"

const (
	EventTwinCreated             = "digitaltwin.created"
	EventTwinStateUpdated        = "digitaltwin.state.updated"
	EventTwinSimulationCompleted = "digitaltwin.simulation.completed"
	EventTwinPlaybackCompleted   = "digitaltwin.playback.completed"
)

// BaseEvent holds common event payload parameters.
type BaseEvent struct {
	EventID   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	PlantID   string    `json:"plant_id"`
	Timestamp time.Time `json:"timestamp"`
	Payload   any       `json:"payload"`
}
