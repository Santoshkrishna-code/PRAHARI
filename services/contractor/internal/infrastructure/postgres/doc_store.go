package postgres

import (
	"context"
	"database/sql"
	"fmt"

	docDomain "prahari/services/contractor/internal/domain/document"
)

// DocStore implements contracts files registry operations.
type DocStore struct {
	db *sql.DB
}

// NewDocStore instantiates DocStore.
func NewDocStore(db *sql.DB) *DocStore {
	return &DocStore{db: db}
}

// Create persists contract metadata.
func (s *DocStore) Create(ctx context.Context, d *docDomain.Document) error {
	query := `INSERT INTO contractor_documents (id, contractor_id, title, storage_path) VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, d.ID, d.ContractorID, d.Title, d.StoragePath)
	return err
}

// FindByContractorID returns documents registered for an onboarding.
func (s *DocStore) FindByContractorID(ctx context.Context, contractorID string) ([]*docDomain.Document, error) {
	query := `SELECT id, contractor_id, title, storage_path FROM contractor_documents WHERE contractor_id = $1`
	rows, err := s.db.QueryContext(ctx, query, contractorID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*docDomain.Document
	for rows.Next() {
		d := &docDomain.Document{}
		err = rows.Scan(&d.ID, &d.ContractorID, &d.Title, &d.StoragePath)
		if err != nil {
			return nil, err
		}
		list = append(list, d)
	}
	return list, nil
}
