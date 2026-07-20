package evaluation

import "time"

// AccuracyReport tracks evaluated inference metrics on Edge devices.
type AccuracyReport struct {
	ID            string    `json:"id"`
	ModelID       string    `json:"model_id"`
	F1Score       float64   `json:"f1_score"`
	PrecisionRate float64   `json:"precision_rate"`
	RecallRate    float64   `json:"recall_rate"`
	TestedAt      time.Time `json:"tested_at"`
}
