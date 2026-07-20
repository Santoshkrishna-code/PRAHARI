package asset

// RegisterAssetCommand carries payload parameters.
type RegisterAssetCommand struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	SerialNumber string `json:"serial_number"`
	DepartmentID string `json:"department_id"`
	LocationID   string `json:"location_id"`
	CategoryID   string `json:"category_id"`
	TypeID       string `json:"type_id"`
	ModelNumber  string `json:"model_number,omitempty"`
}

// UpdateAssetCommand carries profile edit details.
type UpdateAssetCommand struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// TransitionStatusCommand maps lifecycle transitions parameters.
type TransitionStatusCommand struct {
	AssetID    string `json:"asset_id"`
	TargetCode string `json:"target_code"`
	ActorID    string `json:"actor_id"`
	Reason     string `json:"reason,omitempty"`
}
