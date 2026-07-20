package calibrationschedule

import "time"

// Schedule holds planned or upcoming calibration tasks.
type Schedule struct {
	ID           string    `json:"id"`
	InstrumentID string    `json:"instrument_id"`
	PlanID       string    `json:"plan_id"`
	ScheduledFor time.Time `json:"scheduled_for"`
	CompletedAt  *time.Time `json:"completed_at,omitempty"`
	Status       string    `json:"status"` // PENDING, OVERDUE, COMPLETED, CANCELLED
}
