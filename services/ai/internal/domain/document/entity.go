package document

import "time"

// Doc represents raw documents mapped to search context.
type Doc struct {
	ID        string    `json:"id"`
	SourceID  string    `json:"source_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// Chunk represents segmented sections of a document parsed for vector embedding generation.
type Chunk struct {
	ID         string `json:"id"`
	DocID      string `json:"doc_id"`
	Content    string `json:"content"`
	PageNumber int    `json:"page_number"`
}
