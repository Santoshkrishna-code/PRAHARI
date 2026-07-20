package postgres

import (
	"context"
	"database/sql"
	"fmt"

	courseDomain "prahari/services/training/internal/domain/course"
)

// CourseStore implements training courses catalog database.
type CourseStore struct {
	db *sql.DB
}

// NewCourseStore instantiates CourseStore.
func NewCourseStore(db *sql.DB) *CourseStore {
	return &CourseStore{db: db}
}

// Create persists course details.
func (s *CourseStore) Create(ctx context.Context, c *courseDomain.Course) error {
	query := `INSERT INTO courses (id, course_code, title, duration_hours) VALUES ($1, $2, $3, $4)`
	_, err := s.db.ExecContext(ctx, query, c.ID, c.CourseCode, c.Title, c.DurationHours)
	return err
}

// FindByID returns course details.
func (s *CourseStore) FindByID(ctx context.Context, id string) (*courseDomain.Course, error) {
	query := `SELECT id, course_code, title, duration_hours FROM courses WHERE id = $1`
	c := &courseDomain.Course{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.CourseCode, &c.Title, &c.DurationHours)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("course not found: %s", id)
		}
		return nil, err
	}
	return c, nil
}
