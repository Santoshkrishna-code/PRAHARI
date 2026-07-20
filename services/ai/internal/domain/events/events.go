package events

import "time"

const (
	EventSummaryGenerated      = "ai.summary.generated"
	EventRecommendationCreated = "ai.recommendation.created"
	EventPredictionCompleted   = "ai.prediction.completed"
	EventDocumentIndexed       = "ai.document.indexed"
)

// BaseEvent holds common event payload parameters.
type BaseEvent struct {
	EventID   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	PlantID   string    `json:"plant_id"`
	Timestamp time.Time `json:"timestamp"`
	Payload   any       `json:"payload"`
}
