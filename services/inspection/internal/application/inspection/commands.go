package inspection

// CreateInspectionCommand carries payload for schedule setups.
type CreateInspectionCommand struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	InspectionType string `json:"inspection_type"`
	InspectorID    string `json:"inspector_id"`
	DepartmentID   string `json:"department_id"`
	AssetID        string `json:"asset_id,omitempty"`
	LinkedPermitID string `json:"linked_permit_id,omitempty"`
}

// UpdateInspectionCommand carries payload details changes.
type UpdateInspectionCommand struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

// TransitionStatusCommand carries lifecycle flow changes parameter targets.
type TransitionStatusCommand struct {
	InspectionID string `json:"inspection_id"`
	TargetCode   string `json:"target_code"`
	ActorID      string `json:"actor_id"`
	Reason       string `json:"reason,omitempty"`
}
