package barrier

import "time"

// BarrierType classifies the function of the barrier in risk prevention/mitigation.
type BarrierType string

const (
	TypePreventive     BarrierType = "PREVENTIVE"
	TypeDetective      BarrierType = "DETECTIVE"
	TypeProtective     BarrierType = "PROTECTIVE"
	TypeMitigative     BarrierType = "MITIGATIVE"
	TypeRecovery       BarrierType = "RECOVERY"
	TypeAdministrative BarrierType = "ADMINISTRATIVE"
	TypeMechanical     BarrierType = "MECHANICAL"
	TypeInstrumented   BarrierType = "INSTRUMENTED"
	TypeHuman          BarrierType = "HUMAN"
	TypeEmergencyResp  BarrierType = "EMERGENCY_RESPONSE"
)

// Barrier represents a process safety protective barrier aggregate root.
type Barrier struct {
	ID                 string      `json:"id"`
	BarrierCode        string      `json:"barrier_code"`
	PlantID            string      `json:"plant_id"`
	UnitID             string      `json:"unit_id"`
	Title              string      `json:"title"`
	Description        string      `json:"description"`
	Type               BarrierType `json:"type"`
	AssetID            string      `json:"asset_id,omitempty"`
	SILLevel           string      `json:"sil_level,omitempty"` // SIL-1, SIL-2, SIL-3, SIL-4
	IsIPL              bool        `json:"is_ipl"`              // Independent Protection Layer
	PFDTarget          float64     `json:"pfd_target"`          // Target Probability of Failure on Demand
	HealthScore        float64     `json:"health_score"`        // 0.0 to 100.0%
	Status             string      `json:"status"`              // Registered, Assigned, Operational, Inspection, Proof Test, Integrity Assessment, Verified, Bypassed, Impaired, Out of Service, Retired
	LastProofTestedAt *time.Time  `json:"last_proof_tested_at,omitempty"`
	NextProofTestDue  *time.Time  `json:"next_proof_test_due,omitempty"`
	CreatedAt          time.Time   `json:"created_at"`
	UpdatedAt          time.Time   `json:"updated_at"`
}
