package search

import (
	"time"
)

// Criteria structures search params.
type Criteria struct {
	BusinessUnitID           string     `json:"business_unit_id"`
	PlantID                  string     `json:"plant_id"`
	DepartmentID             string     `json:"department_id"`
	ESGObjectiveID           string     `json:"esg_objective_id"`
	Framework                string     `json:"framework"`
	ReportingPeriod          string     `json:"reporting_period"`
	EmissionScope            string     `json:"emission_scope"`
	SustainabilityInitiative string     `json:"sustainability_initiative"`
	KPI                      string     `json:"kpi"`
	ComplianceStatus         string     `json:"compliance_status"`
	StartDate                *time.Time `json:"start_date"`
	EndDate                  *time.Time `json:"end_date"`
	FreeText                 string     `json:"free_text"`
	Limit                    int        `json:"limit"`
	Offset                   int        `json:"offset"`
}
