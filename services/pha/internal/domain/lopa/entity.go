package lopa

import "time"

// Analysis represents a Layer of Protection Analysis (LOPA) evaluation.
type Analysis struct {
	ID                 string    `json:"id"`
	StudyID            string    `json:"study_id"`
	ScenarioID         string    `json:"scenario_id"`
	InitiatingEventFreq float64  `json:"initiating_event_freq"` // Events per year
	TolerableTargetFreq float64  `json:"tolerable_target_freq"` // Target risk frequency
	TotalIPLmitigation  float64  `json:"total_ipl_mitigation"`  // Product of IPL PFDs
	MitigatedEventFreq  float64  `json:"mitigated_event_freq"`  // Freq * Total IPL PFD
	RequiredRRF        float64   `json:"required_rrf"`          // Risk Reduction Factor required
	TargetSIL          string    `json:"target_sil"`            // NONE, SIL-1, SIL-2, SIL-3, SIL-4
	CreatedAt          time.Time `json:"created_at"`
}
