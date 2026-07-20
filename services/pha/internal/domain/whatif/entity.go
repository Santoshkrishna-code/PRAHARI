package whatif

import "time"

// Analysis represents a What-If hazard review item.
type Analysis struct {
	ID             string    `json:"id"`
	StudyID        string    `json:"study_id"`
	Question       string    `json:"question"`       // E.g. "What if the cooling water pump fails?"
	Consequence    string    `json:"consequence"`
	Safeguards     string    `json:"safeguards"`
	Recommendation string    `json:"recommendation"`
	CreatedAt      time.Time `json:"created_at"`
}
