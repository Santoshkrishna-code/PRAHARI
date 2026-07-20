package operatorjournal

import "time"

// Journal represents a control room operator's continuous timeline of readings and event logs.
type Journal struct {
	ID          string    `json:"id"`
	ShiftID     string    `json:"shift_id"`
	OperatorID  string    `json:"operator_id"`
	Subject     string    `json:"subject"`
	Detail      string    `json:"detail"`
	LoggedAt    time.Time `json:"logged_at"`
}
