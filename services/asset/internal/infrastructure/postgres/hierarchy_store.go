package postgres

import (
	"context"
	"database/sql"

	hierarchyDomain "prahari/services/asset/internal/domain/assethierarchy"
)

// HierarchyStore maps tree pathways.
type HierarchyStore struct {
	db *sql.DB
}

// NewHierarchyStore instantiates HierarchyStore.
func NewHierarchyStore(db *sql.DB) *HierarchyStore {
	return &HierarchyStore{db: db}
}

// Create inserts hierarchical relation.
func (s *HierarchyStore) Create(ctx context.Context, h *hierarchyDomain.AssetHierarchy) error {
	query := `INSERT INTO asset_hierarchies (id, parent_asset_id, child_asset_id, relationship_type)
		VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, h.ID, h.ParentAssetID, h.ChildAssetID, h.RelationshipType)
	return err
}

// FindChildren returns matching downstream nodes.
func (s *HierarchyStore) FindChildren(ctx context.Context, parentAssetID string) ([]*hierarchyDomain.AssetHierarchy, error) {
	query := `SELECT id, parent_asset_id, child_asset_id, relationship_type FROM asset_hierarchies WHERE parent_asset_id = $1`
	rows, err := s.db.QueryContext(ctx, query, parentAssetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*hierarchyDomain.AssetHierarchy
	for rows.Next() {
		h := &hierarchyDomain.AssetHierarchy{}
		err = rows.Scan(&h.ID, &h.ParentAssetID, &h.ChildAssetID, &h.RelationshipType)
		if err != nil {
			return nil, err
		}
		list = append(list, h)
	}
	return list, nil
}
