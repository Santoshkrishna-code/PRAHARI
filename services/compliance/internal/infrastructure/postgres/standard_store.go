package postgres

import (
	"context"
	"database/sql"
	"fmt"

	stdDomain "prahari/services/compliance/internal/domain/standard"
)

// StandardStore implements standard frameworks definitions (ISO, Factory Act) catalog.
type StandardStore struct {
	db *sql.DB
}

// NewStandardStore instantiates StandardStore.
func NewStandardStore(db *sql.DB) *StandardStore {
	return &StandardStore{db: db}
}

// Create persists standard rules.
func (s *StandardStore) Create(ctx context.Context, std *stdDomain.Standard) error {
	query := `INSERT INTO standards (id, name, description) VALUES ($1, $2, $3)`
	_, err := s.db.ExecContext(ctx, query, std.ID, std.Name, std.Description)
	return err
}

// FindByID returns standard details.
func (s *StandardStore) FindByID(ctx context.Context, id string) (*stdDomain.Standard, error) {
	query := `SELECT id, name, description FROM standards WHERE id = $1`
	std := &stdDomain.Standard{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&std.ID, &std.Name, &std.Description)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("standard not found: %s", id)
		}
		return nil, err
	}
	return std, nil
}
