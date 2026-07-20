package exposurelimit

import "time"

// Limit defines Occupational Exposure Limits (OEL) for safety monitoring.
type Limit struct {
	ID          string    `json:"id"`
	ChemicalID  string    `json:"chemical_id"`
	LimitType   string    `json:"limit_type"` // TWA, STEL, CEILING, PEL
	ValuePPM    float64   `json:"value_ppm,omitempty"`
	ValueMgM3   float64   `json:"value_mg_m3,omitempty"`
	Source      string    `json:"source"` // OSHA, ACGIH, NIOSH
	EffectiveAt time.Time `json:"effective_at"`
}
