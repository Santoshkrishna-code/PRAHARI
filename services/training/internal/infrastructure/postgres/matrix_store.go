package postgres

import (
	"context"
	"database/sql"

	matrixDomain "prahari/services/training/internal/domain/competencymatrix"
)

// MatrixStore implements competency matrix scopes.
type MatrixStore struct {
	db *sql.DB
}

// NewMatrixStore instantiates MatrixStore.
func NewMatrixStore(db *sql.DB) *MatrixStore {
	return &MatrixStore{db: db}
}

// Create persists matrix definitions.
func (s *MatrixStore) Create(ctx context.Context, cm *matrixDomain.CompetencyMatrix) error {
	query := `INSERT INTO competency_matrix (id, role_id, competency_id) VALUES ($1, $2, $3)`
	_, err := s.db.ExecContext(ctx, query, cm.ID, cm.RoleID, cm.CompetencyID)
	return err
}
