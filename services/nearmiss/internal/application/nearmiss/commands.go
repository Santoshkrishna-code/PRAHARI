package nearmiss

// CreateNearMissCommand carries registry parameters.
type CreateNearMissCommand struct {
	AssetID        string `json:"asset_id,omitempty"`
	ContractorID   string `json:"contractor_id,omitempty"`
	Classification string `json:"classification"`
	SeverityLevel  string `json:"severity_level"`
	DepartmentID   string `json:"department_id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	IsAnonymous    bool   `json:"is_anonymous"`
}

// TransitionStatusCommand carries status updates.
type TransitionStatusCommand struct {
	NearMissID string `json:"near_miss_id"`
	TargetCode string `json:"target_code"`
	ActorID    string `json:"actor_id"`
	Reason     string `json:"reason,omitempty"`
}
