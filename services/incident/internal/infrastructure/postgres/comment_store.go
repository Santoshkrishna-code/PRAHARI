package postgres

import (
	"context"
	"database/sql"
	"fmt"

	commentDomain "prahari/services/incident/internal/domain/comment"
)

// CommentStore implements the comment persistence adapter against PostgreSQL.
type CommentStore struct {
	db *sql.DB
}

// NewCommentStore constructs a CommentStore.
func NewCommentStore(db *sql.DB) *CommentStore {
	return &CommentStore{db: db}
}

// Create persists a new comment.
func (s *CommentStore) Create(ctx context.Context, c *commentDomain.Comment) error {
	query := `INSERT INTO incident_comments (id, incident_id, author_id, parent_comment_id, body, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.ExecContext(ctx, query, c.ID, c.IncidentID, c.AuthorID, c.ParentCommentID, c.Body, c.CreatedAt, c.UpdatedAt)
	if err != nil {
		return fmt.Errorf("postgres: failed to insert comment: %w", err)
	}
	return nil
}

// FindByIncidentID retrieves all comments for an incident ordered chronologically.
func (s *CommentStore) FindByIncidentID(ctx context.Context, incidentID string) ([]*commentDomain.Comment, error) {
	query := `SELECT id, incident_id, author_id, parent_comment_id, body, created_at, updated_at
		FROM incident_comments WHERE incident_id = $1 AND is_deleted = false ORDER BY created_at ASC`
	rows, err := s.db.QueryContext(ctx, query, incidentID)
	if err != nil {
		return nil, fmt.Errorf("postgres: failed to list comments: %w", err)
	}
	defer rows.Close()

	var comments []*commentDomain.Comment
	for rows.Next() {
		c := &commentDomain.Comment{}
		if err := rows.Scan(&c.ID, &c.IncidentID, &c.AuthorID, &c.ParentCommentID, &c.Body, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, fmt.Errorf("postgres: failed to scan comment: %w", err)
		}
		comments = append(comments, c)
	}
	return comments, nil
}
