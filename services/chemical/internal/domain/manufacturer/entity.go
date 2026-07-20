package manufacturer

import "time"

// Manufacturer represents an external entity that produces a chemical.
type Manufacturer struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	ContactInfo string    `json:"contact_info"`
	CreatedAt   time.Time `json:"created_at"`
}
