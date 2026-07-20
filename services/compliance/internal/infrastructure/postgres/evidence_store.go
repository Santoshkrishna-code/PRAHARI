package postgres

import (
	"context"
	"database/sql"

	evidenceDomain "prahari/services/compliance/internal/domain/evidence"
)

// EvidenceStore implements compliance documentation upload files reference tags database.
type EvidenceStore struct {
	db *sql.DB
}

// NewEvidenceStore instantiates EvidenceStore.
func NewEvidenceStore(db *sql.DB) *EvidenceStore {
	return &EvidenceStore{db: db}
}

// Create persists upload document proof tag.
func (s *EvidenceStore) Create(ctx context.Context, e *evidenceDomain.Evidence) error {
	query := `INSERT INTO evidence (id, obligation_id, uploaded_by_id, storage_path, collected_at)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, e.ID, e.ObligationID, e.UploadedByID, e.StoragePath, e.CollectedAt)
	return err
}

// FindByObligationID returns upload files details list.
func (s *EvidenceStore) FindByObligationID(ctx context.Context, obligationID string) ([]*evidenceDomain.Evidence, error) {
	query := `SELECT id, obligation_id, uploaded_by_id, storage_path, collected_at FROM evidence WHERE obligation_id = $1`
	rows, err := s.db.QueryContext(ctx, query, obligationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*evidenceDomain.Evidence
	for rows.Next() {
		e := &evidenceDomain.Evidence{}
		err = rows.Scan(&e.ID, &e.ObligationID, &e.UploadedByID, &e.StoragePath, &e.CollectedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, e)
	}
	return list, nil
}
