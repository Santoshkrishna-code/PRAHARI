package events

import "time"

const (
	EventVisionDetectionCreated  = "vision.detection.created"
	EventVisionAlertTriggered    = "vision.alert.triggered"
	EventVisionInferenceCompleted = "vision.inference.completed"
	EventVisionModelDeployed     = "vision.model.deployed"
	EventVisionCameraOffline     = "vision.camera.offline"
)

// BaseEvent holds common event payload parameters.
type BaseEvent struct {
	EventID   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	PlantID   string    `json:"plant_id"`
	Timestamp time.Time `json:"timestamp"`
	Payload   any       `json:"payload"`
}
