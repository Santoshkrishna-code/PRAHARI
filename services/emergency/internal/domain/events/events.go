package events

import "time"

const (
	EventEmergencyDeclared        = "emergency.declared"
	EventEmergencyResponseActivate = "emergency.response.activated"
	EventEvacuationStarted        = "evacuation.started"
	EventResourceDeployed         = "resource.deployed"
	EventEmergencyStabilized      = "emergency.stabilized"
	EventRecoveryStarted          = "recovery.started"
	EventEmergencyClosed          = "emergency.closed"
)

// BaseEvent holds common event payload parameters.
type BaseEvent struct {
	EventID   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	PlantID   string    `json:"plant_id"`
	Timestamp time.Time `json:"timestamp"`
	Payload   any       `json:"payload"`
}
