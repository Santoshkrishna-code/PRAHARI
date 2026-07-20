package postgres

import (
	"context"
	"database/sql"

	findingDomain "prahari/services/compliance/internal/domain/finding"
)

// FindingStore implements non-compliant EHS violations audit checklist gaps database.
type FindingStore struct {
	db *sql.DB
}

// NewFindingStore instantiates FindingStore.
func NewFindingStore(db *sql.DB) *FindingStore {
	return &FindingStore{db: db}
}

// Create persists violations details.
func (s *FindingStore) Create(ctx context.Context, f *findingDomain.Finding) error {
	query := `INSERT INTO findings (id, compliance_id, severity, description) VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, f.ID, f.ComplianceID, f.Severity, f.Description)
	return err
}

// FindByComplianceID returns violations list.
func (s *FindingStore) FindByComplianceID(ctx context.Context, complianceID string) ([]*findingDomain.Finding, error) {
	query := `SELECT id, compliance_id, severity, description FROM findings WHERE compliance_id = $1`
	rows, err := s.db.QueryContext(ctx, query, complianceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*findingDomain.Finding
	for rows.Next() {
		f := &findingDomain.Finding{}
		err = rows.Scan(&f.ID, &f.ComplianceID, &f.Severity, &f.Description)
		if err != nil {
			return nil, err
		}
		list = append(list, f)
	}
	return list, nil
}
