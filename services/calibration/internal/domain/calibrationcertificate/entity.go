package calibrationcertificate

import "time"

// Certificate represents a generated ISO 17025 compliant calibration document.
type Certificate struct {
	ID             string    `json:"id"`
	CalibrationID  string    `json:"calibration_id"`
	CertificateNo  string    `json:"certificate_no"`
	IssuedDate     time.Time `json:"issued_date"`
	ExpiryDate     time.Time `json:"expiry_date"`
	DocumentDocRef string    `json:"document_doc_ref"` // reference to Document Management
}
