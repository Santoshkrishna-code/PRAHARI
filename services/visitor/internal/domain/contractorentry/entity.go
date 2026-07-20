package contractorentry

import "time"

// Record maps contractor visitor profiles to Contractor Management services.
type Record struct {
	ID           string    `json:"id"`
	VisitID      string    `json:"visit_id"`
	ContractorID string    `json:"contractor_id"`
	WorkOrderRef string    `json:"work_order_ref,omitempty"`
	PermitRef    string    `json:"permit_ref,omitempty"`
	Verified     bool      `json:"verified"`
	CreatedAt    time.Time `json:"created_at"`
}
