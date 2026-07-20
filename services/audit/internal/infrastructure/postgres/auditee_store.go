package postgres

import (
	"context"
	"database/sql"

	auditeeDomain "prahari/services/audit/internal/domain/auditee"
)

// AuditeeStore implements assigned auditees database.
type AuditeeStore struct {
	db *sql.DB
}

// NewAuditeeStore instantiates AuditeeStore.
func NewAuditeeStore(db *sql.DB) *AuditeeStore {
	return &AuditeeStore{db: db}
}

// Create persists auditee records.
func (s *AuditeeStore) Create(ctx context.Context, a *auditeeDomain.Auditee) error {
	query := `INSERT INTO auditees (id, audit_id, user_id) VALUES ($1, $2, $3)`
	_, err := s.db.ExecContext(ctx, query, a.ID, a.AuditID, a.UserID)
	return err
}

// FindByAuditID returns auditees.
func (s *AuditeeStore) FindByAuditID(ctx context.Context, auditID string) ([]*auditeeDomain.Auditee, error) {
	query := `SELECT id, audit_id, user_id FROM auditees WHERE audit_id = $1`
	rows, err := s.db.QueryContext(ctx, query, auditID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*auditeeDomain.Auditee
	for rows.Next() {
		a := &auditeeDomain.Auditee{}
		err = rows.Scan(&a.ID, &a.AuditID, &a.UserID)
		if err != nil {
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil
}
