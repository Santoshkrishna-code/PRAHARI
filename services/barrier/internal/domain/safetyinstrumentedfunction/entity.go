package safetyinstrumentedfunction

import "time"

// SIF represents a Safety Instrumented Function per IEC 61511.
type SIF struct {
	ID            string    `json:"id"`
	BarrierID     string    `json:"barrier_id"`
	SIFNumber     string    `json:"sif_number"`
	TargetSIL     string    `json:"target_sil"`     // SIL-1, SIL-2, SIL-3, SIL-4
	AchievedSIL   string    `json:"achieved_sil"`   // Validated SIL
	SpuriousTrip  float64   `json:"spurious_trip_rate"`
	ProofInterval string    `json:"proof_interval"` // E.g., "12 Months"
	CreatedAt     time.Time `json:"created_at"`
}
