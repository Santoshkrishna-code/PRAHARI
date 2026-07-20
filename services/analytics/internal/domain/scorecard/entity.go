package scorecard

// Card represents calculated monthly/weekly performance card scores.
type Card struct {
	ID             string  `json:"id"`
	PlantID        string  `json:"plant_id"`
	Period         string  `json:"period"` // YYYY-MM
	SafetyScore    float64 `json:"safety_score"`
	ESGScore       float64 `json:"esg_score"`
	UtilitiesScore float64 `json:"utilities_score"`
}
