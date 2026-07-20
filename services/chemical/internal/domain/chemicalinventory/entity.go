package chemicalinventory

import "time"

// Inventory represents the current calculated stock level of a chemical at a specific storage area.
type Inventory struct {
	ID            string    `json:"id"`
	ChemicalID    string    `json:"chemical_id"`
	StorageAreaID string    `json:"storage_area_id"`
	CurrentQty    float64   `json:"current_qty"` // in kg or liters
	ReservedQty   float64   `json:"reserved_qty"`
	UnitOfMeasure string    `json:"unit_of_measure"`
	UpdatedAt     time.Time `json:"updated_at"`
}
