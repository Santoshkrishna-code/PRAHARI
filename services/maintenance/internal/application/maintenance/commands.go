package maintenance

// CreateMaintenanceCommand carries input parameters.
type CreateMaintenanceCommand struct {
	AssetID         string   `json:"asset_id"`
	MaintenanceType string   `json:"maintenance_type"`
	Priority        string   `json:"priority"`
	DepartmentID    string   `json:"department_id"`
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	EstimatedCost   float64  `json:"estimated_cost"`
}

// TransitionStatusCommand carries status updates.
type TransitionStatusCommand struct {
	MaintenanceID string `json:"maintenance_id"`
	TargetCode    string `json:"target_code"`
	ActorID       string `json:"actor_id"`
	Reason        string `json:"reason,omitempty"`
}
