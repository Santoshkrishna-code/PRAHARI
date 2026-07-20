package benchmark

// Comparison represents a calculated metrics benchmarking index between plants.
type Comparison struct {
	ID            string  `json:"id"`
	IndustryAvg   float64 `json:"industry_avg"`
	TargetPlantID string  `json:"target_plant_id"`
	PlantVal      float64 `json:"plant_val"`
	MetricKey     string  `json:"metric_key"`
}
