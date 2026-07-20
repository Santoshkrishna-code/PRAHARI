package lockbox

import "time"

// GroupLockbox represents a group lockbox holding keys to multiple physical isolation locks.
type GroupLockbox struct {
	ID          string    `json:"id"`
	PlantID     string    `json:"plant_id"`
	BoxNumber   string    `json:"box_number"`
	Location    string    `json:"location"`
	MasterKeyID string    `json:"master_key_id"`
	CreatedAt   time.Time `json:"created_at"`
}
