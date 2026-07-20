package batch

import "time"

// Batch represents a single manufactured batch of a chemical.
type Batch struct {
	ID             string    `json:"id"`
	ChemicalID     string    `json:"chemical_id"`
	BatchNumber    string    `json:"batch_number"`
	ManufactureDate time.Time `json:"manufacture_date"`
	ExpiryDate     time.Time `json:"expiry_date"`
	ManufacturerID string    `json:"manufacturer_id"`
	CertOfAnalysis string    `json:"cert_of_analysis,omitempty"` // URL to certificate
	CreatedAt      time.Time `json:"created_at"`
}
