package supplier

import "time"

// Supplier represents an external vender that supplies/distributes a chemical.
type Supplier struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	ContactInfo string    `json:"contact_info"`
	CreatedAt   time.Time `json:"created_at"`
}
