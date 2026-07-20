package effectivenessreview

import "time"

// Review checks if the implemented corrective actions have successfully prevented recurrence after a specified run duration.
type Review struct {
	ID          string    `json:"id"`
	CapaID      string    `json:"capa_id"`
	ReviewedBy  string    `json:"reviewed_by"`
	ReviewedAt  time.Time `json:"reviewed_at"`
	Effective   bool      `json:"effective"`
	Notes       string    `json:"notes"`
	NextCheckAt *time.Time `json:"next_check_at,omitempty"`
}
