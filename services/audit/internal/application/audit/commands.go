package audit

// CreateAuditCommand carries audit parameters.
type CreateAuditCommand struct {
	AssetID      string `json:"asset_id,omitempty"`
	DepartmentID string `json:"department_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

// TransitionStatusCommand carries status updates.
type TransitionStatusCommand struct {
	AuditID    string `json:"audit_id"`
	TargetCode string `json:"target_code"`
	ActorID    string `json:"actor_id"`
	Reason     string `json:"reason,omitempty"`
}
