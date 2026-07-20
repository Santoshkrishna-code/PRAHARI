package riskreview

import "time"

// Review links MOC to fresh Process Hazard Analysis (PHA) or Bow-Tie Risk Assessment.
type Review struct {
	ID                string    `json:"id"`
	ChangeRequestID   string    `json:"change_request_id"`
	RiskAssessmentID  string    `json:"risk_assessment_id,omitempty"` // From Risk Assessment Management Service
	PreChangeRisk     string    `json:"pre_change_risk"`               // LOW, MEDIUM, HIGH, CRITICAL
	PostChangeRisk    string    `json:"post_change_risk"`              // ALARP, LOW, MEDIUM
	MitigationsReqd   string    `json:"mitigations_reqd"`
	RiskManagerID     string    `json:"risk_manager_id"`
	Status            string    `json:"status"` // APPROVED, REJECTED
	ReviewedAt        time.Time `json:"reviewed_at"`
}
