package alert

import (
	"errors"
	"time"
)

// Alert records utility usage violations or high loads.
type Alert struct {
	ID             string    `json:"id" db:"id"`
	PlantID        string    `json:"plant_id" db:"plant_id"`
	MeterID        string    `json:"meter_id" db:"meter_id"`
	AlertLevel     string    `json:"alert_level" db:"alert_level"` // "CRITICAL", "WARNING", "INFO"
	TriggerMessage string    `json:"trigger_message" db:"trigger_message"`
	IsResolved     bool      `json:"is_resolved" db:"is_resolved"`
	TriggeredAt    time.Time `json:"triggered_at" db:"triggered_at"`
	ResolvedAt     time.Time `json:"resolved_at" db:"resolved_at"`
}

// Validate checks fields.
func (a *Alert) Validate() error {
	if a.PlantID == "" {
		return errors.New("plant ID reference is required")
	}
	return nil
}
