package ppeinspection

import "time"

// Record tracks cleaning, safety check, and fitment inspections of active PPE.
type Record struct {
	ID          string    `json:"id"`
	ItemID      string    `json:"item_id"`
	InspectedBy string    `json:"inspected_by"`
	InspectedAt time.Time `json:"inspected_at"`
	Result      string    `json:"result"` // PASS, FAIL
	Findings    string    `json:"findings"`
}
