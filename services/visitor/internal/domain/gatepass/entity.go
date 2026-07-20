package gatepass

import "time"

// Pass represents an authorized gate entry pass with QR code configuration.
type Pass struct {
	ID         string     `json:"id"`
	VisitID    string     `json:"visit_id"`
	PassNumber string     `json:"pass_number"`
	QRCodeURL  string     `json:"qr_code_url"`
	IssuedBy   string     `json:"issued_by"`
	IssuedAt   time.Time  `json:"issued_at"`
	ValidUntil time.Time  `json:"valid_until"`
}
