package search

// Criteria defines multi-dimensional search parameters for admin hierarchies.
type Criteria struct {
	TenantID string `json:"tenant_id,omitempty"`
	Query    string `json:"query,omitempty"`
	Limit    int    `json:"limit,omitempty"`
	Offset   int    `json:"offset,omitempty"`
}
