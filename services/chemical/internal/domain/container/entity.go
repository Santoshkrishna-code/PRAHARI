package container

import "time"

// Container represents a physical tracked container (bottle, drum, cylinder) of a chemical.
type Container struct {
	ID            string     `json:"id"`
	ChemicalID    string     `json:"chemical_id"`
	BatchID       string     `json:"batch_id,omitempty"`
	Barcode       string     `json:"barcode"` // QR / Barcode tag
	StorageAreaID string     `json:"storage_area_id"`
	Capacity      float64    `json:"capacity"`
	CurrentVolume float64    `json:"current_volume"`
	UnitOfMeasure string     `json:"unit_of_measure"`
	Status        string     `json:"status"` // RECEIVED, STORED, ISSUED, IN_USE, DISPOSED, QUARANTINED
	ExpiryDate    *time.Time `json:"expiry_date,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
}
