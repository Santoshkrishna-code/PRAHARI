package search

import "time"

// Criteria defines parameters for multi-faceted water profile and consumption search.
type Criteria struct {
	OrganizationID      string     `json:"organization_id,omitempty"`
	PlantID             string     `json:"plant_id,omitempty"`
	DepartmentID        string     `json:"department_id,omitempty"`
	WaterSource         string     `json:"water_source,omitempty"`
	ReservoirID         string     `json:"reservoir_id,omitempty"`
	DistributionZone    string     `json:"distribution_zone,omitempty"`
	FlowMeterID         string     `json:"flow_meter_id,omitempty"`
	TreatmentPlantID    string     `json:"treatment_plant_id,omitempty"`
	ReportingPeriod     string     `json:"reporting_period,omitempty"`
	ConservationProgram string     `json:"conservation_program,omitempty"`
	StartDate           *time.Time `json:"start_date,omitempty"`
	EndDate             *time.Time `json:"end_date,omitempty"`
	Query               string     `json:"query,omitempty"`
	Limit               int        `json:"limit,omitempty"`
	Offset              int        `json:"offset,omitempty"`
}
