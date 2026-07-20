package postgres

import (
	"context"
	"database/sql"
	"fmt"

	typeDomain "prahari/services/asset/internal/domain/assettype"
)

// TypeStore implements types references mappings.
type TypeStore struct {
	db *sql.DB
}

// NewTypeStore instantiates a TypeStore.
func NewTypeStore(db *sql.DB) *TypeStore {
	return &TypeStore{db: db}
}

// Create inserts type.
func (s *TypeStore) Create(ctx context.Context, at *typeDomain.AssetType) error {
	query := `INSERT INTO asset_types (id, category_id, name, description, is_active)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, at.ID, at.CategoryID, at.Name, at.Description, at.IsActive)
	return err
}

// FindByID returns type.
func (s *TypeStore) FindByID(ctx context.Context, id string) (*typeDomain.AssetType, error) {
	query := `SELECT id, category_id, name, description, is_active FROM asset_types WHERE id = $1`
	at := &typeDomain.AssetType{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&at.ID, &at.CategoryID, &at.Name, &at.Description, &at.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("asset type not found: %s", id)
		}
		return nil, err
	}
	return at, nil
}
