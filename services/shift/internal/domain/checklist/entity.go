package checklist

import "time"

// Item represents a safety / operational check item to be verified at shift start / mid / end.
type Item struct {
	ID          string     `json:"id"`
	ShiftID     string     `json:"shift_id"`
	CheckName   string     `json:"check_name"`
	Completed   bool       `json:"completed"`
	CompletedBy string     `json:"completed_by,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}
