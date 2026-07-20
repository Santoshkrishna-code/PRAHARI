package induction

import "time"

// Verification tracks the safety induction status and validity for visitors.
type Verification struct {
	ID          string    `json:"id"`
	VisitorID   string    `json:"visitor_id"`
	InductionType string  `json:"induction_type"` // Standard Plant, Contractor Safety, Control Room
	CompletedAt time.Time `json:"completed_at"`
	ExpiresAt   time.Time `json:"expires_at"`
	Score       float64   `json:"score"`
}
