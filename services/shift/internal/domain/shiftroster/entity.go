package shiftroster

import "time"

// Roster defines a plan of rotations, detailing which crews work which shifts over a date range.
type Roster struct {
	ID          string    `json:"id"`
	PlantID     string    `json:"plant_id"`
	Name        string    `json:"name"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CalendarID  string    `json:"calendar_id"`
	IsPublished bool      `json:"is_published"`
	CreatedAt   time.Time `json:"created_at"`
}
