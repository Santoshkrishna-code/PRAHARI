package purchase

import "time"

// Order represents a purchase order or request for a chemical.
type Order struct {
	ID          string    `json:"id"`
	PlantID     string    `json:"plant_id"`
	ChemicalID  string    `json:"chemical_id"`
	SupplierID  string    `json:"supplier_id"`
	QtyOrdered  float64   `json:"qty_ordered"`
	UnitOfMeasure string  `json:"unit_of_measure"`
	RequestedBy string    `json:"requested_by"`
	Status      string    `json:"status"` // PENDING, APPROVED, REJECTED, ORDERED, DELIVERED
	CreatedAt   time.Time `json:"created_at"`
}
