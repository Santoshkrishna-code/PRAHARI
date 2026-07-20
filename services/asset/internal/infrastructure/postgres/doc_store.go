package postgres

import (
	"context"
	"database/sql"

	docDomain "prahari/services/asset/internal/domain/document"
)

// DocStore implements drawings references queries.
type DocStore struct {
	db *sql.DB
}

// NewDocStore instantiates DocStore.
func NewDocStore(db *sql.DB) *DocStore {
	return &DocStore{db: db}
}

// Create inserts metadata records.
func (s *DocStore) Create(ctx context.Context, d *docDomain.Document) error {
	query := `INSERT INTO asset_documents (id, asset_id, title, doc_type, storage_path)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, d.ID, d.AssetID, d.Title, d.DocType, d.StoragePath)
	return err
}

// FindByAssetID returns drawings list.
func (s *DocStore) FindByAssetID(ctx context.Context, assetID string) ([]*docDomain.Document, error) {
	query := `SELECT id, asset_id, title, doc_type, storage_path FROM asset_documents WHERE asset_id = $1`
	rows, err := s.db.QueryContext(ctx, query, assetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*docDomain.Document
	for rows.Next() {
		d := &docDomain.Document{}
		err = rows.Scan(&d.ID, &d.AssetID, &d.Title, &d.DocType, &d.StoragePath)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}
