package returnpkg

import "time"

// Transaction represents chemical container return logs.
type Transaction struct {
	ID          string    `json:"id"`
	ContainerID string    `json:"container_id"`
	ReturnedBy  string    `json:"returned_by"`
	ReturnedAt  time.Time `json:"returned_at"`
	QtyReturned float64   `json:"qty_returned"`
	Condition   string    `json:"condition"` // E.g., INTACT, LEAKING, SEALED
}
