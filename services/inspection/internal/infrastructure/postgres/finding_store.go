package postgres

import (
	"context"
	"database/sql"

	findingDomain "prahari/services/inspection/internal/domain/finding"
)

// FindingStore implements failed check logs.
type FindingStore struct {
	db *sql.DB
}

// NewFindingStore instantiates FindingStore.
func NewFindingStore(db *sql.DB) *FindingStore {
	return &FindingStore{db: db}
}

// Create inserts finding logs.
func (s *FindingStore) Create(ctx context.Context, f *findingDomain.Finding) error {
	query := `INSERT INTO inspection_findings (id, inspection_id, checklist_item_id, description, severity, priority, status, identified_by, identified_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := s.db.ExecContext(ctx, query,
		f.ID, f.InspectionID, f.ChecklistItemID, f.Description, f.Severity, f.Priority, f.Status, f.IdentifiedBy, f.IdentifiedAt,
	)
	return err
}

// FindByID returns a finding.
func (s *FindingStore) FindByID(ctx context.Context, id string) (*findingDomain.Finding, error) {
	query := `SELECT id, inspection_id, checklist_item_id, description, severity, priority, status, identified_by, identified_at
		FROM inspection_findings WHERE id = $1`
	f := &findingDomain.Finding{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&f.ID, &f.InspectionID, &f.ChecklistItemID, &f.Description, &f.Severity, &f.Priority, &f.Status, &f.IdentifiedBy, &f.IdentifiedAt)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// FindByInspectionID returns findings on an inspection.
func (s *FindingStore) FindByInspectionID(ctx context.Context, inspectionID string) ([]*findingDomain.Finding, error) {
	query := `SELECT id, inspection_id, checklist_item_id, description, severity, priority, status, identified_by, identified_at
		FROM inspection_findings WHERE inspection_id = $1`
	rows, err := s.db.QueryContext(ctx, query, inspectionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var findings []*findingDomain.Finding
	for rows.Next() {
		f := &findingDomain.Finding{}
		err = rows.Scan(&f.ID, &f.InspectionID, &f.ChecklistItemID, &f.Description, &f.Severity, &f.Priority, &f.Status, &f.IdentifiedBy, &f.IdentifiedAt)
		if err != nil {
			return nil, err
		}
		findings = append(findings, f)
	}
	return findings, nil
}
