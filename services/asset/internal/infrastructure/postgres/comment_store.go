package postgres

import (
	"context"
	"database/sql"

	commentDomain "prahari/services/asset/internal/domain/comment"
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
	query := `INSERT INTO asset_comments (id, asset_id, author_id, parent_comment_id, body, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.ExecContext(ctx, query, c.ID, c.AssetID, c.AuthorID, c.ParentCommentID, c.Body, c.CreatedAt, c.UpdatedAt)
	return err
}

// FindByAssetID returns comment lists.
func (s *CommentStore) FindByAssetID(ctx context.Context, assetID string) ([]*commentDomain.Comment, error) {
	query := `SELECT id, asset_id, author_id, parent_comment_id, body, created_at, updated_at
		FROM asset_comments WHERE asset_id = $1 AND is_deleted = false ORDER BY created_at ASC`
	rows, err := s.db.QueryContext(ctx, query, assetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*commentDomain.Comment
	for rows.Next() {
		c := &commentDomain.Comment{}
		err = rows.Scan(&c.ID, &c.AssetID, &c.AuthorID, &c.ParentCommentID, &c.Body, &c.CreatedAt, &c.UpdatedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, c)
	}
	return list, nil
}
