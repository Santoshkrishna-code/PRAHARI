package traceability

import "time"

// Record maps calibration certificate standards back to NIST, NPL, or international primary standard reference certificates.
type Record struct {
	ID                 string    `json:"id"`
	ReferenceStandardID string    `json:"reference_standard_id"`
	PrimaryCertNo      string    `json:"primary_cert_no"`
	CalibratedByBody   string    `json:"calibrated_by_body"`
	VerifiedAt         time.Time `json:"verified_at"`
}
