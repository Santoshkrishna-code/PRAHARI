package search

// Criteria defines parameters to query vectorized document chunks.
type Criteria struct {
	Query string `json:"query"`
	Limit int    `json:"limit,omitempty"`
}
