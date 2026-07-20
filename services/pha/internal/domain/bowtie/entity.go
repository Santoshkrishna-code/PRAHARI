package bowtie

import "time"

// Analysis represents a BowTie risk visualization model (Threats -> Top Event -> Consequences).
type Analysis struct {
	ID                string    `json:"id"`
	StudyID           string    `json:"study_id"`
	TopEvent          string    `json:"top_event"`
	HazardDescription string    `json:"hazard_description"`
	PreventiveBarriers string   `json:"preventive_barriers"` // Threats & left side barriers
	MitigationBarriers string   `json:"mitigation_barriers"` // Right side barriers
	CreatedAt         time.Time `json:"created_at"`
}
