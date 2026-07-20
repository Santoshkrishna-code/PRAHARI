package ppeitem

import "time"

// Item represents a specific serialized physical instance of a PPE (e.g. barcode / RFID code hard hat).
type Item struct {
	ID          string     `json:"id"`
	PPEID       string     `json:"ppe_id"`
	SerialNumber string    `json:"serial_number"`
	RFIDCode    string     `json:"rfid_code,omitempty"`
	Barcode     string     `json:"barcode,omitempty"`
	ManufactureDate time.Time `json:"manufacture_date"`
	ExpiryDate  time.Time  `json:"expiry_date"`
	Status      string     `json:"status"` // Available, Reserved, Issued, In Use, Inspection, Returned, Maintenance, Expired, Disposed
	IssuedTo    string     `json:"issued_to,omitempty"` // UserID / VisitorID
	LastInspectedAt *time.Time `json:"last_inspected_at,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
}
