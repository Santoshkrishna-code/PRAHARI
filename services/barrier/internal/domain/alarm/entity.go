package alarm

import "time"

// Alarm represents a safety-critical alarm barrier per ISA 18.2.
type Alarm struct {
	ID           string    `json:"id"`
	BarrierID    string    `json:"barrier_id"`
	TagNumber    string    `json:"tag_number"`
	Priority     string    `json:"priority"` // LOW, MEDIUM, HIGH, EMERGENCY
	SetpointVal  float64   `json:"setpoint_val"`
	ResponseTime float64   `json:"response_time_min"`
	CreatedAt    time.Time `json:"created_at"`
}
