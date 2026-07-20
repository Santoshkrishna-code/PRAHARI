package postgres

import (
	"context"
	"database/sql"

	specDomain "prahari/services/asset/internal/domain/specification"
)

// SpecStore implements technical dynamic specifications fields queries.
type SpecStore struct {
	db *sql.DB
}

// NewSpecStore instantiates SpecStore.
func NewSpecStore(db *sql.DB) *SpecStore {
	return &SpecStore{db: db}
}

// Create inserts specifications.
func (s *SpecStore) Create(ctx context.Context, sp *specDomain.Specification) error {
	query := `INSERT INTO asset_specifications (id, asset_id, attributes) VALUES ($1, $2, $3)`
	_, err := s.db.ExecContext(ctx, query, sp.ID, sp.AssetID, sp.Attributes)
	return err
}

// FindByAssetID returns specifications map.
func (s *SpecStore) FindByAssetID(ctx context.Context, assetID string) (*specDomain.Specification, error) {
	query := `SELECT id, asset_id, attributes FROM asset_specifications WHERE asset_id = $1`
	sp := &specDomain.Specification{}
	err := s.db.QueryRowContext(ctx, query, assetID).Scan(&sp.ID, &sp.AssetID, &sp.Attributes)
	if err != nil {
		return nil, err
	}
	return sp, nil
}
