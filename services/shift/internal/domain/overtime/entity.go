package overtime

import "time"

// Record tracks overtime worked on shifts.
type Record struct {
	ID          string    `json:"id"`
	ShiftID     string    `json:"shift_id"`
	UserID      string    `json:"user_id"`
	HoursWorked float64   `json:"hours_worked"`
	Reason      string    `json:"reason"`
	ApprovedBy  string    `json:"approved_by,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}
