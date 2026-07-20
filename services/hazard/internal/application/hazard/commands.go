package hazard

// CreateHazardCommand carries registry parameters.
type CreateHazardCommand struct {
	AssetID      string `json:"asset_id,omitempty"`
	ContractorID string `json:"contractor_id,omitempty"`
	HazardType   string `json:"hazard_type"`
	DepartmentID string `json:"department_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

// TransitionStatusCommand carries status updates.
type TransitionStatusCommand struct {
	HazardID   string `json:"hazard_id"`
	TargetCode string `json:"target_code"`
	ActorID    string `json:"actor_id"`
	Reason     string `json:"reason,omitempty"`
}
