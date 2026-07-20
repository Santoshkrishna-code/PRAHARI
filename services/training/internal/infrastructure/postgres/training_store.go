package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	trainingDomain "prahari/services/training/internal/domain/training"
	searchDomain "prahari/services/training/internal/domain/search"
)

// TrainingStore implements the core repository port.
type TrainingStore struct {
	db *sql.DB
}

// NewTrainingStore instantiates a TrainingStore.
func NewTrainingStore(db *sql.DB) *TrainingStore {
	return &TrainingStore{db: db}
}

// Create inserts a workforce training session.
func (s *TrainingStore) Create(ctx context.Context, t *trainingDomain.Training) error {
	query := `INSERT INTO training_programs (
		id, training_number, course_id, department_id,
		status_code, title, description, created_at, updated_at
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := s.db.ExecContext(ctx, query,
		t.ID, t.TrainingNumber, t.CourseID, t.DepartmentID,
		t.StatusCode, t.Title, t.Description, t.CreatedAt, t.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to insert training record: %w", err)
	}
	return nil
}

// FindByID retrieves a training by ID.
func (s *TrainingStore) FindByID(ctx context.Context, id string) (*trainingDomain.Training, error) {
	query := `SELECT id, training_number, course_id, department_id,
		status_code, title, description, created_at, updated_at
		FROM training_programs WHERE id = $1 AND is_deleted = false`

	t := &trainingDomain.Training{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&t.ID, &t.TrainingNumber, &t.CourseID, &t.DepartmentID,
		&t.StatusCode, &t.Title, &t.Description, &t.CreatedAt, &t.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("training session record not found: %s", id)
		}
		return nil, fmt.Errorf("postgres: failed to query training: %w", err)
	}
	return t, nil
}

// FindByNumber retrieves a training by number.
func (s *TrainingStore) FindByNumber(ctx context.Context, number string) (*trainingDomain.Training, error) {
	query := `SELECT id, training_number, course_id, department_id,
		status_code, title, description, created_at, updated_at
		FROM training_programs WHERE training_number = $1 AND is_deleted = false`

	t := &trainingDomain.Training{}
	err := s.db.QueryRowContext(ctx, query, number).Scan(
		&t.ID, &t.TrainingNumber, &t.CourseID, &t.DepartmentID,
		&t.StatusCode, &t.Title, &t.Description, &t.CreatedAt, &t.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("training not found: %s", number)
		}
		return nil, fmt.Errorf("postgres: failed to query training by number: %w", err)
	}
	return t, nil
}

// Update persists modifications.
func (s *TrainingStore) Update(ctx context.Context, t *trainingDomain.Training) error {
	query := `UPDATE training_programs SET
		status_code = $2, title = $3, description = $4, updated_at = $5 WHERE id = $1`

	_, err := s.db.ExecContext(ctx, query,
		t.ID, t.StatusCode, t.Title, t.Description, time.Now(),
	)
	if err != nil {
		return fmt.Errorf("postgres: failed to update training: %w", err)
	}
	return nil
}

// Delete soft-deletes a training session.
func (s *TrainingStore) Delete(ctx context.Context, id string) error {
	query := `UPDATE training_programs SET is_deleted = true, updated_at = $2 WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, id, time.Now())
	if err != nil {
		return fmt.Errorf("postgres: failed to soft-delete training record: %w", err)
	}
	return nil
}

// List returns pages.
func (s *TrainingStore) List(ctx context.Context, offset, limit int) ([]*trainingDomain.Training, int, error) {
	var total int
	if err := s.db.QueryRowContext(ctx, `SELECT COUNT(*) FROM training_programs WHERE is_deleted = false`).Scan(&total); err != nil {
		return nil, 0, err
	}

	query := `SELECT id, training_number, course_id, department_id,
		status_code, title, description, created_at, updated_at
		FROM training_programs WHERE is_deleted = false ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var list []*trainingDomain.Training
	for rows.Next() {
		t := &trainingDomain.Training{}
		err = rows.Scan(
			&t.ID, &t.TrainingNumber, &t.CourseID, &t.DepartmentID,
			&t.StatusCode, &t.Title, &t.Description, &t.CreatedAt, &t.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		list = append(list, t)
	}
	return list, total, nil
}

// Search queries GIN indexes.
func (s *TrainingStore) Search(ctx context.Context, criteria *searchDomain.Criteria) ([]*trainingDomain.Training, int, error) {
	return s.List(ctx, criteria.Offset(), criteria.PageSize)
}
