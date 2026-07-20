package graph

// Edge represents a relationship between two equipment nodes.
type Edge struct {
	ID           string `json:"id"`
	TwinID       string `json:"twin_id"`
	FromNodeID   string `json:"from_node_id"`
	ToNodeID     string `json:"to_node_id"`
	RelationType string `json:"relation_type"` // FEEDS_INTO, CONTROLS, MONITORS, PROTECTS
}
