package postgres

import (
	"context"
	"database/sql"

	programDomain "prahari/services/audit/internal/domain/auditprogram"
)

// ProgramStore implements annual scheduling tracks database.
type ProgramStore struct {
	db *sql.DB
}

// NewProgramStore instantiates ProgramStore.
func NewProgramStore(db *sql.DB) *ProgramStore {
	return &ProgramStore{db: db}
}

// Create persists program tracker.
func (s *ProgramStore) Create(ctx context.Context, ap *programDomain.AuditProgram) error {
	query := `INSERT INTO audit_programs (id, name, description) VALUES ($1, $2, $3)`
	_, err := s.db.ExecContext(ctx, query, ap.ID, ap.Name, ap.Description)
	return err
}
