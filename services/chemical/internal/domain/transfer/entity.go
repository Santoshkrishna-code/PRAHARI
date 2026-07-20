package transfer

import "time"

// Transaction represents chemical transfer logs between different storage locations.
type Transaction struct {
	ID          string    `json:"id"`
	ContainerID string    `json:"container_id"`
	FromAreaID  string    `json:"from_area_id"`
	ToAreaID    string    `json:"to_area_id"`
	TransferredBy string  `json:"transferred_by"`
	TransferredAt time.Time `json:"transferred_at"`
}
