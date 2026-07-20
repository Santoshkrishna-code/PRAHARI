package search

import (
	"time"
)

// Criteria structures search params.
type Criteria struct {
	Organization    string     `json:"organization"`
	PlantID         string     `json:"plant_id"`
	DepartmentID    string     `json:"department_id"`
	ProductionLine  string     `json:"production_line"`
	AssetID         string     `json:"asset_id"`
	UtilityMeterID  string     `json:"utility_meter_id"`
	EnergySourceID  string     `json:"energy_source_id"`
	ReportingPeriod string     `json:"reporting_period"`
	Tariff          string     `json:"tariff"`
	EnergyTargetID  string     `json:"energy_target_id"`
	StartDate       *time.Time `json:"start_date"`
	EndDate         *time.Time `json:"end_date"`
	FreeText        string     `json:"free_text"`
	Limit           int        `json:"limit"`
	Offset          int        `json:"offset"`
}
