package impactassessment

import "time"

// Assessment evaluates multi-disciplinary impacts of a proposed change.
type Assessment struct {
	ID                  string    `json:"id"`
	ChangeRequestID     string    `json:"change_request_id"`
	SafetyImpact        bool      `json:"safety_impact"`
	EnvironmentalImpact bool      `json:"environmental_impact"`
	QualityImpact       bool      `json:"quality_impact"`
	ReliabilityImpact   bool      `json:"reliability_impact"`
	CybersecurityImpact bool      `json:"cybersecurity_impact"`
	RegulatoryImpact    bool      `json:"regulatory_impact"`
	PAndIDImpact        bool      `json:"p_and_id_impact"` // Piping & Instrumentation Diagram
	HAZOPRequired       bool      `json:"hazop_required"`
	SummaryNotes        string    `json:"summary_notes"`
	AssessedBy          string    `json:"assessed_by"`
	AssessedAt          time.Time `json:"assessed_at"`
}
