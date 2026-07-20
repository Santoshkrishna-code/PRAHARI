package receiving

import "time"

// Delivery represents a chemical batch delivery transaction.
type Delivery struct {
	ID          string    `json:"id"`
	OrderID     string    `json:"order_id,omitempty"`
	PlantID     string    `json:"plant_id"`
	ChemicalID  string    `json:"chemical_id"`
	SupplierID  string    `json:"supplier_id"`
	QtyReceived float64   `json:"qty_received"`
	UnitOfMeasure string  `json:"unit_of_measure"`
	ReceivedBy  string    `json:"received_by"`
	ReceivedAt  time.Time `json:"received_at"`
	LotNumber   string    `json:"lot_number"`
}
