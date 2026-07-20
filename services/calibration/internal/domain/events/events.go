package events

import "time"

const (
	EventCalibrationScheduled  = "calibration.scheduled"
	EventCalibrationStarted    = "calibration.started"
	EventCalibrationCompleted  = "calibration.completed"
	EventCalibrationFailed     = "calibration.failed"
	EventCertificateGenerated  = "certificate.generated"
	EventInstrumentOutOfTolerance = "instrument.out_of_tolerance"
)

// BaseEvent holds common event payload parameters.
type BaseEvent struct {
	EventID   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	PlantID   string    `json:"plant_id"`
	Timestamp time.Time `json:"timestamp"`
	Payload   any       `json:"payload"`
}
