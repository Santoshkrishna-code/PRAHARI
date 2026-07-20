package postgres

import (
	"context"
	"database/sql"

	programDomain "prahari/services/training/internal/domain/trainingprogram"
)

// ProgramStore implements structured tracks catalog.
type ProgramStore struct {
	db *sql.DB
}

// NewProgramStore instantiates ProgramStore.
func NewProgramStore(db *sql.DB) *ProgramStore {
	return &ProgramStore{db: db}
}

// Create persists program name.
func (s *ProgramStore) Create(ctx context.Context, tp *programDomain.TrainingProgram) error {
	query := `INSERT INTO courses (id, name, description) VALUES ($1, $2, $3)`
	_, err := s.db.ExecContext(ctx, query, tp.ID, tp.Name, tp.Description)
	return err
}
