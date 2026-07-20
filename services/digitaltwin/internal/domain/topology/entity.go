package topology

// Node represents a node in the plant topology graph.
type Node struct {
	ID         string `json:"id"`
	TwinID     string `json:"twin_id"`
	FacilityID string `json:"facility_id"`
	Label      string `json:"label"` // E.g., Distillation Column DC-101
	NodeType   string `json:"node_type"` // PROCESS, UTILITY, SAFETY, STORAGE
}
