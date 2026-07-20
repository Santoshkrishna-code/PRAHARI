package changerequest

import "time"

// ChangeCategory defines whether a change is permanent, temporary, or emergency.
type ChangeCategory string

const (
	CategoryPermanent    ChangeCategory = "PERMANENT"
	CategoryTemporary    ChangeCategory = "TEMPORARY"
	CategoryEmergency    ChangeCategory = "EMERGENCY"
	CategoryOrganizational ChangeCategory = "ORGANIZATIONAL"
)

// Request represents an enterprise Management of Change (MOC) request.
type Request struct {
	ID              string         `json:"id"`
	MOCNumber       string         `json:"moc_number"`
	PlantID         string         `json:"plant_id"`
	DepartmentID    string         `json:"department_id"`
	Title           string         `json:"title"`
	Description     string         `json:"description"`
	ReasonForChange string         `json:"reason_for_change"`
	Category        ChangeCategory `json:"category"`
	ChangeType      string         `json:"change_type"` // Process, Mechanical, Electrical, Instrumentation, Software, etc.
	TargetAssetID   string         `json:"target_asset_id,omitempty"`
	RiskLevel       string         `json:"risk_level"` // LOW, MEDIUM, HIGH, CRITICAL
	Status          string         `json:"status"`     // Draft, Impact Assessment, Technical Review, Risk Review, Safety Review, Approval, Implementation, Verification, Closeout, Rejected, Cancelled, Rolled Back
	RequesterID     string         `json:"requester_id"`
	TargetDate      time.Time      `json:"target_date"`
	ExpiryDate      *time.Time     `json:"expiry_date,omitempty"` // For temporary changes
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
}
