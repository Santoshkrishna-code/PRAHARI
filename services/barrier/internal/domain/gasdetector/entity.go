package gasdetector

import "time"

// Detector represents a toxic or flammable gas detector barrier.
type Detector struct {
	ID             string    `json:"id"`
	BarrierID      string    `json:"barrier_id"`
	TagNumber      string    `json:"tag_number"`
	GasType        string    `json:"gas_type"` // H2S, CH4, CO, LEL
	AlarmThreshold float64   `json:"alarm_threshold"`
	LocationCode   string    `json:"location_code"`
	CreatedAt      time.Time `json:"created_at"`
}
