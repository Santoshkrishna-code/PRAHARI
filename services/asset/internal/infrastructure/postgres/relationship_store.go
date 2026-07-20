package postgres

import (
	"context"
	"database/sql"

	relationshipDomain "prahari/services/asset/internal/domain/relationship"
)

// RelationshipStore maps dependency flows.
type RelationshipStore struct {
	db *sql.DB
}

// NewRelationshipStore instantiates RelationshipStore.
func NewRelationshipStore(db *sql.DB) *RelationshipStore {
	return &RelationshipStore{db: db}
}

// Create persists relationship links.
func (s *RelationshipStore) Create(ctx context.Context, r *relationshipDomain.Relationship) error {
	query := `INSERT INTO asset_relationships (id, source_asset_id, target_asset_id, dependency_type)
		VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, r.ID, r.SourceAssetID, r.TargetAssetID, r.DependencyType)
	return err
}

// FindDependencies returns downstream matches.
func (s *RelationshipStore) FindDependencies(ctx context.Context, sourceAssetID string) ([]*relationshipDomain.Relationship, error) {
	query := `SELECT id, source_asset_id, target_asset_id, dependency_type FROM asset_relationships WHERE source_asset_id = $1`
	rows, err := s.db.QueryContext(ctx, query, sourceAssetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*relationshipDomain.Relationship
	for rows.Next() {
		r := &relationshipDomain.Relationship{}
		err = rows.Scan(&r.ID, &r.SourceAssetID, &r.TargetAssetID, &r.DependencyType)
		if err != nil {
			return nil, err
		}
		list = append(list, r)
	}
	return list, nil
}
