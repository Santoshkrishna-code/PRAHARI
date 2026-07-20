package compliance

// CreateComplianceCommand carries register definitions.
type CreateComplianceCommand struct {
	AssetID      string `json:"asset_id,omitempty"`
	DepartmentID string `json:"department_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

// TransitionStatusCommand carries status updates.
type TransitionStatusCommand struct {
	ComplianceID string `json:"compliance_id"`
	TargetCode   string `json:"target_code"`
	ActorID      string `json:"actor_id"`
	Reason       string `json:"reason,omitempty"`
}
