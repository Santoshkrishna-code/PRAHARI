package verification

import "time"

// Record tracks Pre-Startup Safety Review (PSSR) and post-implementation validation.
type Record struct {
	ID              string    `json:"id"`
	ChangeRequestID string    `json:"change_request_id"`
	PSSRCompleted   bool      `json:"pssr_completed"`
	TrainingVerified bool     `json:"training_verified"`
	DocsUpdated     bool      `json:"docs_updated"`
	VerifiedBy      string    `json:"verified_by"`
	Status          string    `json:"status"` // VERIFIED_PASSED, VERIFIED_FAILED
	Comments        string    `json:"comments"`
	VerifiedAt      time.Time `json:"verified_at"`
}
