package ppeissue

import "time"

// Record tracks physical issuance of protective equipment to an employee, contractor, or visitor.
type Record struct {
	ID          string    `json:"id"`
	ItemID      string    `json:"item_id"`
	IssuedToType string   `json:"issued_to_type"` // EMPLOYEE, CONTRACTOR, VISITOR
	IssuedToID  string    `json:"issued_to_id"`
	IssuedBy    string    `json:"issued_by"`
	IssuedAt    time.Time `json:"issued_at"`
	ExpectedReturn time.Time `json:"expected_return"`
}
