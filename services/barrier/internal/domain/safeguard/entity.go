package safeguard

import "time"

// Safeguard represents a general protective feature mapped to a barrier.
type Safeguard struct {
	ID           string    `json:"id"`
	BarrierID    string    `json:"barrier_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	SafeguardType string   `json:"safeguard_type"`
	CreatedAt    time.Time `json:"created_at"`
}
