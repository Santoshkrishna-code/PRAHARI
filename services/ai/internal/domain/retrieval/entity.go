package retrieval

// Result represents returned search hits with similarity confidence metrics.
type Result struct {
	ChunkID    string  `json:"chunk_id"`
	DocID      string  `json:"doc_id"`
	Text       string  `json:"text"`
	Confidence float64 `json:"confidence"`
}
