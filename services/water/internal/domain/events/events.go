package events

import "time"

const (
	EventWaterProfileCreated   = "water.profile.created"
	EventFlowmeterReadingRec   = "flowmeter.reading.recorded"
	EventWaterForecastGen      = "water.forecast.generated"
	EventWaterTargetExceeded   = "water.target.exceeded"
	EventWaterLeakDetected     = "water.leak.detected"
	EventWaterRecycled         = "water.recycled"
	EventWaterDeleted          = "water.deleted"
)

// BaseEvent holds common event payload parameters.
type BaseEvent struct {
	EventID   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	PlantID   string    `json:"plant_id"`
	Timestamp time.Time `json:"timestamp"`
	Payload   any       `json:"payload"`
}
