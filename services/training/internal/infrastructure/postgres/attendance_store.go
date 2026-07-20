package postgres

import (
	"context"
	"database/sql"

	attendanceDomain "prahari/services/training/internal/domain/attendance"
)

// AttendanceStore implements attendance logs database.
type AttendanceStore struct {
	db *sql.DB
}

// NewAttendanceStore instantiates AttendanceStore.
func NewAttendanceStore(db *sql.DB) *AttendanceStore {
	return &AttendanceStore{db: db}
}

// Create persists attendance metrics.
func (s *AttendanceStore) Create(ctx context.Context, a *attendanceDomain.Attendance) error {
	query := `INSERT INTO attendance (id, training_id, trainee_id, attended_date, is_present)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, a.ID, a.TrainingID, a.TraineeID, a.AttendedDate, a.IsPresent)
	return err
}
