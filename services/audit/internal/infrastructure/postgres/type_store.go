package postgres

import (
	"context"
	"database/sql"
	"fmt"

	typeDomain "prahari/services/audit/internal/domain/audittype"
)

// TypeStore implements audit types catalog classifications.
type TypeStore struct {
	db *sql.DB
}

// NewTypeStore instantiates TypeStore.
func NewTypeStore(db *sql.DB) *TypeStore {
	return &TypeStore{db: db}
}

// Create persists audit classification type code.
func (s *TypeStore) Create(ctx context.Context, at *typeDomain.AuditType) error {
	query := `INSERT INTO audit_types (id, code, name, description) VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, at.ID, at.Code, at.Name, at.Description)
	return err
}

// FindByID returns audit type details.
func (s *TypeStore) FindByID(ctx context.Context, id string) (*typeDomain.AuditType, error) {
	query := `SELECT id, code, name, description FROM audit_types WHERE id = $1`
	at := &typeDomain.AuditType{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&at.ID, &at.Code, &at.Name, &at.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("audit type not found: %s", id)
		}
		return nil, err
	}
	return at, nil
}
