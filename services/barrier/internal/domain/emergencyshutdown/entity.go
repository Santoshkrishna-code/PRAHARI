package emergencyshutdown

import "time"

// System represents an Emergency Shutdown (ESD) valve or trip system.
type System struct {
	ID             string    `json:"id"`
	BarrierID      string    `json:"barrier_id"`
	SystemName     string    `json:"system_name"`
	ValveTagNumber string    `json:"valve_tag_number"`
	StrokeTimeSec  float64   `json:"stroke_time_sec"`
	CreatedAt      time.Time `json:"created_at"`
}
