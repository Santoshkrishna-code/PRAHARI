package shift

import "time"

// Shift represents an active or scheduled work shift in a plant.
type Shift struct {
	ID              string     `json:"id"`
	ShiftName       string     `json:"shift_name"` // E.g., Day Shift, Night Shift, Swing Shift
	PlantID         string     `json:"plant_id"`
	UnitID          string     `json:"unit_id"`
	SupervisorID    string     `json:"supervisor_id"`
	ScheduledStart  time.Time  `json:"scheduled_start"`
	ScheduledEnd    time.Time  `json:"scheduled_end"`
	ActualStart     *time.Time `json:"actual_start,omitempty"`
	ActualEnd       *time.Time `json:"actual_end,omitempty"`
	Status          string     `json:"status"` // Scheduled, Crew Assigned, Shift Started, Operational, Handover Initiated, Handover Accepted, Shift Closed, Cancelled
	HandoverID      string     `json:"handover_id,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}
