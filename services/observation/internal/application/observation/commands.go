package observation

// CreateObservationCommand carries registry parameters.
type CreateObservationCommand struct {
	AssetID         string `json:"asset_id,omitempty"`
	ContractorID    string `json:"contractor_id,omitempty"`
	ObservationType string `json:"observation_type"`
	DepartmentID    string `json:"department_id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
}

// TransitionStatusCommand carries status updates.
type TransitionStatusCommand struct {
	ObservationID string `json:"observation_id"`
	TargetCode    string `json:"target_code"`
	ActorID       string `json:"actor_id"`
	Reason        string `json:"reason,omitempty"`
}
