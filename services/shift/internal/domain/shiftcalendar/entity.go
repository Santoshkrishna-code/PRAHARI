package shiftcalendar

import "time"

// Calendar represents a shift scheduling pattern (e.g. Dupont shift schedule, 2-2-3 schedule).
type Calendar struct {
	ID          string    `json:"id"`
	PlantID     string    `json:"plant_id"`
	PatternName string    `json:"pattern_name"`
	CycleDays   int       `json:"cycle_days"`
	Description string    `json:"description"`
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"created_at"`
}
