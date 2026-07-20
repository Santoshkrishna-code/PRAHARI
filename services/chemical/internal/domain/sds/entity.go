package sds

import "time"

// SDS represents the header definition of a Safety Data Sheet.
type SDS struct {
	ID           string    `json:"id"`
	ChemicalID   string    `json:"chemical_id"`
	Version      string    `json:"version"`
	Manufacturer string    `json:"manufacturer"`
	PublishDate  time.Time `json:"publish_date"`
	DocumentURL  string    `json:"document_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
