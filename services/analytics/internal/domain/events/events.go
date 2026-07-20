package events

import "time"

const (
	EventReportGenerated      = "report.generated"
	EventDashboardUpdated     = "dashboard.updated"
	EventKPIThresholdExceeded = "kpi.threshold.exceeded"
	EventAnalyticsCompleted   = "analytics.completed"
)

// BaseEvent holds common event payload parameters.
type BaseEvent struct {
	EventID   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	PlantID   string    `json:"plant_id"`
	Timestamp time.Time `json:"timestamp"`
	Payload   any       `json:"payload"`
}
