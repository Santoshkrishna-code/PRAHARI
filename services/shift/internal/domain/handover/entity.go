package handover

import "time"

// Handover tracks the transition from the outgoing crew to the incoming crew, capturing open safety/operational metrics.
type Handover struct {
	ID                 string     `json:"id"`
	OutgoingShiftID    string     `json:"outgoing_shift_id"`
	IncomingShiftID    string     `json:"incoming_shift_id"`
	OutgoingLeadID     string     `json:"outgoing_lead_id"`
	IncomingLeadID     string     `json:"incoming_lead_id"`
	OpenPermitIDs      string     `json:"open_permit_ids"` // Comma-separated or descriptive text
	ActiveMaintenance  string     `json:"active_maintenance"`
	SafetyIncidents    string     `json:"safety_incidents"`
	OperationalContinuityNotes string `json:"operational_continuity_notes"`
	InitiatedAt        time.Time  `json:"initiated_at"`
	AcceptedAt         *time.Time `json:"accepted_at,omitempty"`
	Status             string     `json:"status"` // PENDING, ACCEPTED, REJECTED
}
