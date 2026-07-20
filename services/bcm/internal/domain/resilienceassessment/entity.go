package resilienceassessment

import "time"

// Assessment evaluates holistic organizational resilience per ISO 22301 standards.
type Assessment struct {
	ID              string    `json:"id"`
	PlantID         string    `json:"plant_id"`
	BusinessUnit    string    `json:"business_unit"`
	ResilienceIndex float64   `json:"resilience_index_pct"` // 0 to 100%
	ISO22301Status  string    `json:"iso22301_status"`      // COMPLIANT, MINOR_NON_CONFORMANCE, NON_COMPLIANT
	AssessedBy      string    `json:"assessed_by"`
	AssessedAt      time.Time `json:"assessed_at"`
}
