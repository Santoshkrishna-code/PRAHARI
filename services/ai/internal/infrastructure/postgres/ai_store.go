package postgres

import (
	"context"
	"database/sql"
	"time"

	"prahari/services/ai/internal/domain/conversation"
	"prahari/services/ai/internal/domain/document"
	"prahari/services/ai/internal/domain/prediction"
	"prahari/services/ai/internal/domain/recommendation"
	"prahari/services/ai/internal/domain/search"
	"prahari/services/ai/internal/domain/summarization"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) SaveThread(ctx context.Context, t *conversation.Thread) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO conversations (id, user_id, plant_id, created_at) VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, t.ID, t.UserID, t.PlantID, t.CreatedAt)
	return err
}

func (s *Store) SaveMessage(ctx context.Context, m *conversation.Message) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO messages (id, thread_id, role, content, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, m.ID, m.ThreadID, m.Role, m.Content, m.CreatedAt)
	return err
}

func (s *Store) GetMessagesByThread(ctx context.Context, threadID string) ([]*conversation.Message, error) {
	if s.db == nil {
		return []*conversation.Message{
			{ID: "msg-1", ThreadID: threadID, Role: "USER", Content: "Hello RCA assistant", CreatedAt: time.Now()},
		}, nil
	}
	query := `SELECT id, thread_id, role, content, created_at FROM messages WHERE thread_id = $1 ORDER BY created_at ASC`
	rows, err := s.db.QueryContext(ctx, query, threadID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*conversation.Message
	for rows.Next() {
		var m conversation.Message
		if err := rows.Scan(&m.ID, &m.ThreadID, &m.Role, &m.Content, &m.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, &m)
	}
	return result, nil
}

func (s *Store) SaveDocument(ctx context.Context, d *document.Doc) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO knowledge_documents (id, source_id, title, content, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, d.ID, d.SourceID, d.Title, d.Content, d.CreatedAt)
	return err
}

func (s *Store) SaveChunk(ctx context.Context, c *document.Chunk) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO document_chunks (id, doc_id, content, page_number) VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, c.ID, c.DocID, c.Content, c.PageNumber)
	return err
}

func (s *Store) SaveRecommendation(ctx context.Context, r *recommendation.Recommendation) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO recommendations (id, plant_id, type, source_id, content, created_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, r.ID, r.PlantID, r.Type, r.SourceID, r.Content, r.CreatedAt)
	return err
}

func (s *Store) SaveSummary(ctx context.Context, sum *summarization.Summary) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO summaries (id, source_id, original, condensed, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, sum.ID, sum.SourceID, sum.Original, sum.Condensed, sum.CreatedAt)
	return err
}

func (s *Store) SaveForecast(ctx context.Context, f *prediction.Forecast) error {
	if s.db == nil {
		return nil
	}
	query := `INSERT INTO predictions (id, plant_id, target_topic, probability, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, f.ID, f.PlantID, f.TargetTopic, f.Probability, f.CreatedAt)
	return err
}

func (s *Store) SearchDocuments(ctx context.Context, criteria *search.Criteria) ([]*document.Doc, int64, error) {
	if s.db == nil {
		mockDoc := &document.Doc{ID: "doc-001", SourceID: "s-1", Title: "LOTO Safety Manual", Content: "Lock out tag out criteria rules", CreatedAt: time.Now()}
		return []*document.Doc{mockDoc}, 1, nil
	}
	query := `SELECT id, source_id, title, content, created_at FROM knowledge_documents WHERE title ILIKE $1`
	rows, err := s.db.QueryContext(ctx, query, "%"+criteria.Query+"%")
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var result []*document.Doc
	for rows.Next() {
		var d document.Doc
		if err := rows.Scan(&d.ID, &d.SourceID, &d.Title, &d.Content, &d.CreatedAt); err != nil {
			return nil, 0, err
		}
		result = append(result, &d)
	}
	return result, int64(len(result)), nil
}

func (s *Store) GetDocumentByID(ctx context.Context, id string) (*document.Doc, error) {
	return &document.Doc{ID: id, SourceID: "s-01", Title: "LOTO Safety Manual", Content: "Lock out tag out criteria rules", CreatedAt: time.Now()}, nil
}
