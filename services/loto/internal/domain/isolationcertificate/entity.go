package isolationcertificate

import "time"

// Certificate represents a permit-linked authorization for executing the isolation plan.
type Certificate struct {
	ID              string     `json:"id"`
	PlanID          string     `json:"plan_id"`
	PermitID        string     `json:"permit_id,omitempty"` // Link to Permit-to-Work
	IssuerID        string     `json:"issuer_id"`
	ReceiverID      string     `json:"receiver_id"`
	Status          string     `json:"status"` // Planned, Isolation Approved, Locks Applied, Tags Applied, Zero Energy Verified, Maintenance In Progress, Restoration Planned, Locks Removed, Returned To Service, Closed, Cancelled, Failed Verification
	VerifiedAt      *time.Time `json:"verified_at,omitempty"`
	RestoredAt      *time.Time `json:"restored_at,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}
