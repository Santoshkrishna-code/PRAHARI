package contractor

import (
	"time"
)

// RegisterContractorCommand carries registry parameters.
type RegisterContractorCommand struct {
	CompanyName      string    `json:"company_name"`
	TaxID            string    `json:"tax_id"`
	DepartmentID     string    `json:"department_id"`
	InsuranceExpiry  time.Time `json:"insurance_expiry"`
}

// TransitionStatusCommand carries status updates.
type TransitionStatusCommand struct {
	ContractorID string `json:"contractor_id"`
	TargetCode   string `json:"target_code"`
	ActorID      string `json:"actor_id"`
	Reason       string `json:"reason,omitempty"`
}
