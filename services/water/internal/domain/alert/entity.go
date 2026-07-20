package alert

import "time"

// Alert represents a water system alert.
type Alert struct {
	ID             string    `json:"id"`
	PlantID        string    `json:"plant_id"`
	MeterID        string    `json:"meter_id,omitempty"`
	AlertLevel     string    `json:"alert_level"`
	TriggerMessage string    `json:"trigger_message"`
	IsResolved     bool      `json:"is_resolved"`
	TriggeredAt    time.Time `json:"triggered_at"`
	ResolvedAt     time.Time `json:"resolved_at,omitempty"`
}
