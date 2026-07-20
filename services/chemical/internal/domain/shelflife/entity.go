package shelflife

// Rule defines product shelf-life storage parameters.
type Rule struct {
	ID          string  `json:"id"`
	ChemicalID  string  `json:"chemical_id"`
	LifeDays    int     `json:"life_days"`
	AlertThresholdDays int `json:"alert_threshold_days"` // Days before expiry to trigger alert
}
