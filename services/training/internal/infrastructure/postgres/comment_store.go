package postgres

import (
	"context"
	"database/sql"

	commentDomain "prahari/services/training/internal/domain/comment"
)

// CommentStore implements threaded comments database.
type CommentStore struct {
	db *sql.DB
}

// NewCommentStore instantiates CommentStore.
func NewCommentStore(db *sql.DB) *CommentStore {
	return &CommentStore{db: db}
}

// Create persists comments notes.
func (s *CommentStore) Create(ctx context.Context, c *commentDomain.Comment) error {
	query := `INSERT INTO comments (id, training_id, author_id, parent_comment_id, body, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.ExecContext(ctx, query, c.ID, c.TrainingID, c.AuthorID, c.ParentCommentID, c.Body, c.CreatedAt, c.UpdatedAt)
	return err
}

// FindByTrainingID returns comments list.
func (s *CommentStore) FindByTrainingID(ctx context.Context, trainingID string) ([]*commentDomain.Comment, error) {
	query := `SELECT id, training_id, author_id, parent_comment_id, body, created_at, updated_at
		FROM comments WHERE training_id = $1 AND is_deleted = false ORDER BY created_at ASC`
	rows, err := s.db.QueryContext(ctx, query, trainingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*commentDomain.Comment
	for rows.Next() {
		c := &commentDomain.Comment{}
		err = rows.Scan(&c.ID, &c.TrainingID, &c.AuthorID, &c.ParentCommentID, &c.Body, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, c)
	}
	return list, nil
}
