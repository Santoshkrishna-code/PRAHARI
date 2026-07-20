package postgres

import (
	"context"
	"database/sql"

	commentDomain "prahari/services/maintenance/internal/domain/comment"
)

// CommentStore implements profiles walkthrough discussion logs.
type CommentStore struct {
	db *sql.DB
}

// NewCommentStore instantiates CommentStore.
func NewCommentStore(db *sql.DB) *CommentStore {
	return &CommentStore{db: db}
}

// Create persists comments notes.
func (s *CommentStore) Create(ctx context.Context, c *commentDomain.Comment) error {
	query := `INSERT INTO maintenance_comments (id, maintenance_id, author_id, parent_comment_id, body, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.ExecContext(ctx, query, c.ID, c.MaintenanceID, c.AuthorID, c.ParentCommentID, c.Body, c.CreatedAt, c.UpdatedAt)
	return err
}

// FindByMaintenanceID returns comment lists.
func (s *CommentStore) FindByMaintenanceID(ctx context.Context, maintenanceID string) ([]*commentDomain.Comment, error) {
	query := `SELECT id, maintenance_id, author_id, parent_comment_id, body, created_at, updated_at
		FROM maintenance_comments WHERE maintenance_id = $1 AND is_deleted = false ORDER BY created_at ASC`
	rows, err := s.db.QueryContext(ctx, query, maintenanceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*commentDomain.Comment
	for rows.Next() {
		c := &commentDomain.Comment{}
		err = rows.Scan(&c.ID, &c.MaintenanceID, &c.AuthorID, &c.ParentCommentID, &c.Body, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, c)
	}
	return list, nil
}
