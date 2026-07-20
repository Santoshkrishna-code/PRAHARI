package postgres

import (
	"context"
	"database/sql"
	"fmt"
)

// LessonStore implements lessons learned logs storage.
type LessonStore struct {
	db *sql.DB
}

// NewLessonStore instantiates LessonStore.
func NewLessonStore(db *sql.DB) *LessonStore {
	return &LessonStore{db: db}
}

// SaveLesson registers review notes.
func (s *LessonStore) SaveLesson(ctx context.Context, nearmissID, summary string) error {
	query := `INSERT INTO near_miss_lessons (id, near_miss_id, summary) VALUES ($1, $2, $3)`
	_, err := s.db.ExecContext(ctx, query, nearmissID, nearmissID, summary)
	if err != nil {
		return fmt.Errorf("postgres: failed to insert lesson: %w", err)
	}
	return nil
}
