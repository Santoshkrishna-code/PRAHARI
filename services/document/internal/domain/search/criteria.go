package search

import "time"

// Criteria defines multi-dimensional full-text search parameters for controlled documents.
type Criteria struct {
	PlantID      string     `json:"plant_id,omitempty"`
	CategoryID   string     `json:"category_id,omitempty"`
	DocumentType string     `json:"document_type,omitempty"`
	Status       string     `json:"status,omitempty"`
	OwnerID      string     `json:"owner_id,omitempty"`
	OverdueOnly  bool       `json:"overdue_only,omitempty"`
	StartDate    *time.Time `json:"start_date,omitempty"`
	EndDate      *time.Time `json:"end_date,omitempty"`
	Query        string     `json:"query,omitempty"`
	Limit        int        `json:"limit,omitempty"`
	Offset       int        `json:"offset,omitempty"`
}
