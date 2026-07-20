package kpi

import "time"

// KPI represents a Key Performance Indicator target definition and actual result.
type KPI struct {
	ID        string    `json:"id"`
	PlantID   string    `json:"plant_id"`
	Name      string    `json:"name"` // E.g., LTI_free_days
	TargetVal float64   `json:"target_val"`
	ActualVal float64   `json:"actual_val"`
	Status    string    `json:"status"` // ON_TRACK, AT_RISK, CRITICAL
	UpdatedAt time.Time `json:"updated_at"`
}
