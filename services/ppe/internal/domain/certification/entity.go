package certification

import "time"

// Record tracks standards certification (e.g. CE, NIOSH) for specific PPE models.
type Record struct {
	ID                string    `json:"id"`
	PPEID             string    `json:"ppe_id"`
	CertifyingBody    string    `json:"certifying_body"`
	CertificationCode string    `json:"certification_code"`
	IssuedDate        time.Time `json:"issued_date"`
	ExpiryDate        time.Time `json:"expiry_date"`
	DocumentDocRef    string    `json:"document_doc_ref"` // reference to Document Management
}
