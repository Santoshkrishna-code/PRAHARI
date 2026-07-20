package sparepart

import (
	"errors"
)

// SparePart tracks material replacements lists.
type SparePart struct {
	ID            string  `json:"id" db:"id"`
	MaintenanceID string  `json:"maintenance_id" db:"maintenance_id"`
	PartNumber    string  `json:"part_number" db:"part_number"`
	QuantityUsed  int     `json:"quantity_used" db:"quantity_used"`
	UnitCost      float64 `json:"unit_cost" db:"unit_cost"`
}

// Validate checks domain invariants.
func (sp *SparePart) Validate() error {
	if sp.MaintenanceID == "" {
		return errors.New("maintenance ID reference is required")
	}
	if sp.PartNumber == "" {
		return errors.New("part number code is required")
	}
	if sp.QuantityUsed <= 0 {
		return errors.New("quantity used must exceed zero")
	}
	return nil
}

// TotalCost calculated product value sums.
func (sp *SparePart) TotalCost() float64 {
	return float64(sp.QuantityUsed) * sp.UnitCost
}
