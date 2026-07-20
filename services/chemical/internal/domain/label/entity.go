package label

import "time"

// Label represents a printed tag or label identifier containing chemical hazard warnings.
type Label struct {
	ID          string    `json:"id"`
	ChemicalID  string    `json:"chemical_id"`
	NFPAHealth  int       `json:"nfpa_health"` // 0-4
	NFPAFire    int       `json:"nfpa_fire"`   // 0-4
	NFPANoReact int       `json:"nfpa_noreact"` // 0-4
	NFPASpecial string    `json:"nfpa_special,omitempty"`
	LabelFormat string    `json:"label_format"` // GHS, NFPA, CUSTOM
	GeneratedAt time.Time `json:"generated_at"`
	GeneratedBy string    `json:"generated_by"`
}
