package attendance

import "time"

// Record tracks attendance of shift workers.
type Record struct {
	ID         string     `json:"id"`
	ShiftID    string     `json:"shift_id"`
	UserID     string     `json:"user_id"`
	Present    bool       `json:"present"`
	CheckInAt  *time.Time `json:"check_in_at,omitempty"`
	CheckOutAt *time.Time `json:"check_out_at,omitempty"`
}
