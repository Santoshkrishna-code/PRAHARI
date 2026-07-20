package waste

import (
	"errors"
	"time"
)

// SolidWaste tracks generation and recycling/disposal of plant trash.
type SolidWaste struct {
	ID             string    `json:"id" db:"id"`
	PlantID        string    `json:"plant_id" db:"plant_id"`
	WasteCategory  string    `json:"waste_category" db:"waste_category"` // "PLASTIC", "PAPER", "METAL", "ORGANIC", "MIXED"
	WeightKg       float64   `json:"weight_kg" db:"weight_kg"`
	DisposalMethod string    `json:"disposal_method" db:"disposal_method"` // "RECYCLE", "LANDFILL", "INCINERATE"
	DisposalDate   time.Time `json:"disposal_date" db:"disposal_date"`
	VendorName     string    `json:"vendor_name" db:"vendor_name"`
	IsRecycled     bool      `json:"is_recycled" db:"is_recycled"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

// Validate checks waste inputs.
func (w *SolidWaste) Validate() error {
	if w.PlantID == "" {
		return errors.New("plant ID is required")
	}
	if w.WasteCategory == "" {
		return errors.New("waste category type is required")
	}
	return nil
}
