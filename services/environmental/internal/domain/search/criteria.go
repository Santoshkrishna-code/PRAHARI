package search

import (
	"time"
)

// Criteria structures search params.
type Criteria struct {
	SiteID           string     `json:"site_id"`
	PlantID          string     `json:"plant_id"`
	DepartmentID     string     `json:"department_id"`
	AspectCategory   string     `json:"aspect_category"`
	PermitNumber     string     `json:"permit_number"`
	WasteType        string     `json:"waste_type"`
	EmissionSourceID string     `json:"emission_source_id"`
	ProgramID        string     `json:"program_id"`
	LaboratoryID     string     `json:"laboratory_id"`
	ComplianceStatus string     `json:"compliance_status"`
	StartDate        *time.Time `json:"start_date"`
	EndDate          *time.Time `json:"end_date"`
	FreeText         string     `json:"free_text"`
	Limit            int        `json:"limit"`
	Offset           int        `json:"offset"`
}
