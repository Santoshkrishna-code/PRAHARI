package barrierfunction

import "time"

// Function defines the specific safety action expected from a barrier.
type Function struct {
	ID             string    `json:"id"`
	BarrierID      string    `json:"barrier_id"`
	FunctionName   string    `json:"function_name"`
	RequiredAction string    `json:"required_action"`
	ResponseTimeSec float64  `json:"response_time_sec"`
	CreatedAt      time.Time `json:"created_at"`
}
