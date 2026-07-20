package vector

// Data represents generated floating point embeddings mapped to source text chunks.
type Data struct {
	ID        string    `json:"id"`
	ChunkID   string    `json:"chunk_id"`
	Embedding []float32 `json:"embedding"`
}
