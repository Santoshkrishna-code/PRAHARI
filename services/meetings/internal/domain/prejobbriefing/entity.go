package prejobbriefing

import "time"

// Briefing represents a pre-job safety briefing conducted before hazardous work begins,
// typically linked to a Permit-to-Work.
type Briefing struct {
	ID            string    `json:"id"`
	MeetingID     string    `json:"meeting_id"`
	PermitID      string    `json:"permit_id"`
	WorkOrderID   string    `json:"work_order_id,omitempty"`
	HazardsSummary string  `json:"hazards_summary"`
	PPERequired   string    `json:"ppe_required"`
	EmergencyPlan string    `json:"emergency_plan"`
	Acknowledged  bool      `json:"acknowledged"`
	CreatedAt     time.Time `json:"created_at"`
}
