package safetyreview

import "time"

// Review evaluates process safety and occupational health implications of a change.
type Review struct {
	ID                 string    `json:"id"`
	ChangeRequestID    string    `json:"change_request_id"`
	PSMImpactVerified  bool      `json:"psm_impact_verified"`
	OccupationalHealth bool      `json:"occupational_health"`
	EmergencyResponse  bool      `json:"emergency_response"`
	SafetyOfficerID    string    `json:"safety_officer_id"`
	Status             string    `json:"status"` // APPROVED, REJECTED, CHANGES_REQUESTED
	Comments           string    `json:"comments"`
	ReviewedAt         time.Time `json:"reviewed_at"`
}
