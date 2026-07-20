package hazardouswaste

import (
	"errors"
	"time"
)

// HazardousWaste monitors regulated chemical/toxic byproducts.
type HazardousWaste struct {
	ID                 string    `json:"id" db:"id"`
	PlantID            string    `json:"plant_id" db:"plant_id"`
	ChemicalName       string    `json:"chemical_name" db:"chemical_name"`
	UNNumber           string    `json:"un_number" db:"un_number"` // UN Hazmat classification
	WeightKg           float64   `json:"weight_kg" db:"weight_kg"`
	StorageLocation    string    `json:"storage_location" db:"storage_location"`
	ManifestDocumentID string    `json:"manifest_document_id" db:"manifest_document_id"`
	DisposalMethod     string    `json:"disposal_method" db:"disposal_method"`
	AuthorizedVendorID string    `json:"authorized_vendor_id" db:"authorized_vendor_id"`
	DisposalDate       time.Time `json:"disposal_date" db:"disposal_date"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
}

// Validate checks hazmat attributes.
func (h *HazardousWaste) Validate() error {
	if h.PlantID == "" {
		return errors.New("plant ID is required")
	}
	if h.ChemicalName == "" {
		return errors.New("hazardous chemical name is required")
	}
	return nil
}
