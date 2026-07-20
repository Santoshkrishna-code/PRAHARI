package issue

import "time"

// Transaction represents chemical container issuance for job tasks.
type Transaction struct {
	ID          string    `json:"id"`
	ContainerID string    `json:"container_id"`
	IssuedTo    string    `json:"issued_to"`
	IssuedBy    string    `json:"issued_by"`
	IssuedAt    time.Time `json:"issued_at"`
	PermitID    string    `json:"permit_id,omitempty"`
	WorkOrderID string    `json:"work_order_id,omitempty"`
	QtyIssued   float64   `json:"qty_issued"`
}
