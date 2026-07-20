package postgres

import (
	"context"
	"database/sql"

	commentDomain "prahari/services/observation/internal/domain/comment"
)

// CommentStore implements threaded comments notes query.
type CommentStore struct {
	db *sql.DB
}

// NewCommentStore instantiates CommentStore.
func NewCommentStore(db *sql.DB) *CommentStore {
	return &CommentStore{db: db}
}

// Create persists comments notes.
func (s *CommentStore) Create(ctx context.Context, c *commentDomain.Comment) error {
	query := `INSERT INTO comments (id, observation_id, author_id, parent_comment_id, body, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.ExecContext(ctx, query, c.ID, c.ObservationID, c.AuthorID, c.ParentCommentID, c.Body, c.CreatedAt, c.UpdatedAt)
	return err
}

// FindByObservationID returns comments lists.
func (s *CommentStore) FindByObservationID(ctx context.Context, observationID string) ([]*commentDomain.Comment, error) {
	query := `SELECT id, observation_id, author_id, parent_comment_id, body, created_at, updated_at
		FROM comments WHERE observation_id = $1 AND is_deleted = false ORDER BY created_at ASC`
	rows, err := s.db.QueryContext(ctx, query, observationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*commentDomain.Comment
	for rows.Next() {
		c := &commentDomain.Comment{}
		err = rows.Scan(&c.ID, &c.ObservationID, &c.AuthorID, &c.ParentCommentID, &c.Body, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, c)
	}
	return list, nil
}
