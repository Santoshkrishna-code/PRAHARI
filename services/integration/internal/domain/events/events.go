package events

import "time"

const (
	EventIntegrationCompleted   = "integration.completed"
	EventIntegrationFailed      = "integration.failed"
	EventConnectorConnected     = "connector.connected"
	EventConnectorDisconnected  = "connector.disconnected"
	EventMessageTransformed     = "message.transformed"
	EventRetryCompleted         = "retry.completed"
)

// BaseEvent holds common event payload parameters.
type BaseEvent struct {
	EventID   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	PlantID   string    `json:"plant_id"`
	Timestamp time.Time `json:"timestamp"`
	Payload   any       `json:"payload"`
}
