package events

import "time"

const (
	EventPPECreated              = "ppe.created"
	EventPPEIssued               = "ppe.issued"
	EventPPEReturned             = "ppe.returned"
	EventPPEInspected            = "ppe.inspected"
	EventPPEMaintenanceCompleted = "ppe.maintenance.completed"
	EventPPEExpired              = "ppe.expired"
)

// BaseEvent holds common event payload parameters.
type BaseEvent struct {
	EventID   string    `json:"event_id"`
	EventType string    `json:"event_type"`
	PlantID   string    `json:"plant_id"`
	Timestamp time.Time `json:"timestamp"`
	Payload   any       `json:"payload"`
}
