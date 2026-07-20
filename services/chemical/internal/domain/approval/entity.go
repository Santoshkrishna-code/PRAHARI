package approval

import "time"

// Request represents a workflow process to introduce a new chemical to the plant site.
type Request struct {
	ID                  string     `json:"id"`
	PlantID             string     `json:"plant_id"`
	ChemicalName        string     `json:"chemical_name"`
	CASNumber           string     `json:"cas_number"`
	RequestedBy         string     `json:"requested_by"`
	TechnicalReviewerID string     `json:"technical_reviewer_id,omitempty"`
	TechnicalApprovedAt *time.Time `json:"technical_approved_at,omitempty"`
	SafetyReviewerID    string     `json:"safety_reviewer_id,omitempty"`
	SafetyApprovedAt    *time.Time `json:"safety_approved_at,omitempty"`
	EnvReviewerID       string     `json:"env_reviewer_id,omitempty"`
	EnvApprovedAt       *time.Time `json:"env_approved_at,omitempty"`
	Status              string     `json:"status"` // PENDING_REVIEW, TECHNICAL_REVIEW, SAFETY_REVIEW, ENV_REVIEW, APPROVED, REJECTED
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
}
