package outoftolerance

import "time"

// Case tracks investigation into critical out-of-tolerance (OOT) measurements that might compromise plant process safety.
type Case struct {
	ID             string    `json:"id"`
	CalibrationID  string    `json:"calibration_id"`
	ReportedBy     string    `json:"reported_by"`
	ReportedAt     time.Time `json:"reported_at"`
	ImpactAnalysis string    `json:"impact_analysis"`
	RootCause      string    `json:"root_cause,omitempty"`
	Status         string    `json:"status"` // INVESTIGATING, CONCLUDED, ESCALATED
}
