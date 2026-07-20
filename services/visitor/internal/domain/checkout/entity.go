package checkout

import "time"

// Record tracks physical plant check-out event details.
type Record struct {
	ID           string    `json:"id"`
	VisitID      string    `json:"visit_id"`
	CheckOutAt   time.Time `json:"check_out_at"`
	CheckedOutBy string    `json:"checked_out_by"`
	BadgeReturned bool     `json:"badge_returned"`
}
