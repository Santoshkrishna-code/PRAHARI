package inventory

import "time"

// Stock represents current stock levels, safety buffer stock triggers, and warehouse location of PPE models.
type Stock struct {
	ID             string    `json:"id"`
	PPEID          string    `json:"ppe_id"`
	PlantID        string    `json:"plant_id"`
	QuantityOnHand int       `json:"quantity_on_hand"`
	BufferLevel    int       `json:"buffer_level"` // Reorder threshold
	Location       string    `json:"location"`     // Rack A-12, Bin 4, etc.
	UpdatedAt      time.Time `json:"updated_at"`
}
