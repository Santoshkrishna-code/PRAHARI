package calibration

import "time"

// Record tracks a single calibration execution event.
type Record struct {
	ID             string     `json:"id"`
	InstrumentID   string     `json:"instrument_id"`
	CalibratedBy   string     `json:"calibrated_by"`
	CalibratedAt   time.Time  `json:"calibrated_at"`
	Status         string     `json:"status"` // Registered, Scheduled, Calibration Started, Measurement Recorded, Tolerance Verification, Certificate Generated, Approved, Active, Failed, Out of Tolerance, Retired
	Result         string     `json:"result"` // PASS, FAIL, ADJUSTED
	CertificateID  string     `json:"certificate_id,omitempty"`
	ApprovedBy     string     `json:"approved_by,omitempty"`
	ApprovedAt     *time.Time `json:"approved_at,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}
