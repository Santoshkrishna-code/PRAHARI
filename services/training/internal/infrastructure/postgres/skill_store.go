package postgres

import (
	"context"
	"database/sql"

	skillDomain "prahari/services/training/internal/domain/skill"
)

// SkillStore implements skills definitions database.
type SkillStore struct {
	db *sql.DB
}

// NewSkillStore instantiates SkillStore.
func NewSkillStore(db *sql.DB) *SkillStore {
	return &SkillStore{db: db}
}

// Create persists skill attributes.
func (s *SkillStore) Create(ctx context.Context, sk *skillDomain.Skill) error {
	query := `INSERT INTO skills (id, name, description) VALUES ($1, $2, $3)`
	_, err := s.db.ExecContext(ctx, query, sk.ID, sk.Name, sk.Description)
	return err
}
