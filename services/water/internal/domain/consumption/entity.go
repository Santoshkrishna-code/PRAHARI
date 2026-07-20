package consumption

import "time"

// Consumption represents aggregated water usage for a period.
type Consumption struct {
	ID               string    `json:"id"`
	PlantID          string    `json:"plant_id"`
	MeterID          string    `json:"meter_id"`
	PeriodStart      time.Time `json:"period_start"`
	PeriodEnd        time.Time `json:"period_end"`
	ConsumptionKL    float64   `json:"consumption_kl"`
	PeakFlowKLH      float64   `json:"peak_flow_klh"`
	WaterIntensity   float64   `json:"water_intensity"`
	ProductionUnits  float64   `json:"production_units"`
	CostUSD          float64   `json:"cost_usd"`
	CreatedAt        time.Time `json:"created_at"`
}
