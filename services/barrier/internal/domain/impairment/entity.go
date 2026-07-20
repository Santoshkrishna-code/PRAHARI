package impairment

import "time"

// Record tracks barrier degradation or partial loss of protection.
type Record struct {
	ID             string     `json:"id"`
	BarrierID      string     `json:"barrier_id"`
	Reason         string     `json:"reason"`
	CompensatingCtrl string   `json:"compensating_ctrl"`
	ImpairedBy     string     `json:"impaired_by"`
	IsActive       bool       `json:"is_active"`
	ImpairedAt     time.Time  `json:"impaired_at"`
	RestoredAt     *time.Time `json:"restored_at,omitempty"`
}
