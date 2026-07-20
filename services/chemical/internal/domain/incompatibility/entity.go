package incompatibility

// Record log of incompatibility alerts.
type Record struct {
	ID            string `json:"id"`
	ChemicalID    string `json:"chemical_id"`
	OtherChemID   string `json:"other_chem_id"`
	AlertReason   string `json:"alert_reason"`
	SeverityLevel string `json:"severity_level"` // WARNING, CRITICAL
}
