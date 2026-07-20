package assessment

// CreateRiskCommand carries creation inputs.
type CreateRiskCommand struct {
	AssetID      string `json:"asset_id,omitempty"`
	DepartmentID string `json:"department_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

// TransitionStatusCommand carries status updates.
type TransitionStatusCommand struct {
	RiskID     string `json:"risk_id"`
	TargetCode string `json:"target_code"`
	ActorID    string `json:"actor_id"`
	Reason     string `json:"reason,omitempty"`
}
