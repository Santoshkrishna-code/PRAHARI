package postgres

import (
	"context"
	"database/sql"

	evidenceDomain "prahari/services/audit/internal/domain/evidence"
)

// EvidenceStore implements audit evidence documentation references uploads.
type EvidenceStore struct {
	db *sql.DB
}

// NewEvidenceStore instantiates EvidenceStore.
func NewEvidenceStore(db *sql.DB) *EvidenceStore {
	return &EvidenceStore{db: db}
}

// Create persists upload document tags.
func (s *EvidenceStore) Create(ctx context.Context, e *evidenceDomain.Evidence) error {
	query := `INSERT INTO evidence (id, audit_id, uploaded_by_id, storage_path, collected_at)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, e.ID, e.AuditID, e.UploadedByID, e.StoragePath, e.CollectedAt)
	return err
}

// FindByAuditID returns evidence list.
func (s *EvidenceStore) FindByAuditID(ctx context.Context, auditID string) ([]*evidenceDomain.Evidence, error) {
	query := `SELECT id, audit_id, uploaded_by_id, storage_path, collected_at FROM evidence WHERE audit_id = $1`
	rows, err := s.db.QueryContext(ctx, query, auditID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*evidenceDomain.Evidence
	for rows.Next() {
		e := &evidenceDomain.Evidence{}
		err = rows.Scan(&e.ID, &e.AuditID, &e.UploadedByID, &e.StoragePath, &e.CollectedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, e)
	}
	return list, nil
}
