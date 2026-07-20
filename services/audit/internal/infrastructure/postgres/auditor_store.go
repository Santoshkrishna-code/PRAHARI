package postgres

import (
	"context"
	"database/sql"

	auditorDomain "prahari/services/audit/internal/domain/auditor"
)

// AuditorStore implements assigned auditors mapping databases.
type AuditorStore struct {
	db *sql.DB
}

// NewAuditorStore instantiates AuditorStore.
func NewAuditorStore(db *sql.DB) *AuditorStore {
	return &AuditorStore{db: db}
}

// Create persists auditor credentials constraints.
func (s *AuditorStore) Create(ctx context.Context, a *auditorDomain.Auditor) error {
	query := `INSERT INTO auditors (id, audit_id, user_id, role) VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, a.ID, a.AuditID, a.UserID, a.Role)
	return err
}

// FindByAuditID returns auditor roster.
func (s *AuditorStore) FindByAuditID(ctx context.Context, auditID string) ([]*auditorDomain.Auditor, error) {
	query := `SELECT id, audit_id, user_id, role FROM auditors WHERE audit_id = $1`
	rows, err := s.db.QueryContext(ctx, query, auditID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*auditorDomain.Auditor
	for rows.Next() {
		a := &auditorDomain.Auditor{}
		err = rows.Scan(&a.ID, &a.AuditID, &a.UserID, &a.Role)
		if err != nil {
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil
}
