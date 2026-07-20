package postgres

import (
	"context"
	"database/sql"

	findingDomain "prahari/services/audit/internal/domain/finding"
)

// FindingStore implements observations and NCR corrective CAPA tasks database.
type FindingStore struct {
	db *sql.DB
}

// NewFindingStore instantiates FindingStore.
func NewFindingStore(db *sql.DB) *FindingStore {
	return &FindingStore{db: db}
}

// Create persists audit findings gaps.
func (s *FindingStore) Create(ctx context.Context, f *findingDomain.Finding) error {
	query := `INSERT INTO findings (id, audit_id, finding_type, description) VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, f.ID, f.AuditID, f.FindingType, f.Description)
	return err
}

// FindByAuditID returns findings.
func (s *FindingStore) FindByAuditID(ctx context.Context, auditID string) ([]*findingDomain.Finding, error) {
	query := `SELECT id, audit_id, finding_type, description FROM findings WHERE audit_id = $1`
	rows, err := s.db.QueryContext(ctx, query, auditID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*findingDomain.Finding
	for rows.Next() {
		f := &findingDomain.Finding{}
		err = rows.Scan(&f.ID, &f.AuditID, &f.FindingType, &f.Description)
		if err != nil {
			return nil, err
		}
		list = append(list, f)
	}
	return list, nil
}
