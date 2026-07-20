package postgres

import (
	"context"
	"database/sql"
	"fmt"

	scheduleDomain "prahari/services/inspection/internal/domain/schedule"
)

// ScheduleStore implements recurring scheduler settings.
type ScheduleStore struct {
	db *sql.DB
}

// NewScheduleStore instantiates ScheduleStore.
func NewScheduleStore(db *sql.DB) *ScheduleStore {
	return &ScheduleStore{db: db}
}

// Create persists schedule setup.
func (s *ScheduleStore) Create(ctx context.Context, sch *scheduleDomain.Schedule) error {
	query := `INSERT INTO inspection_schedule (id, template_id, frequency, inspector_id, department_id, last_execution_date, next_execution_date, is_active)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := s.db.ExecContext(ctx, query,
		sch.ID, sch.TemplateID, sch.Frequency, sch.InspectorID, sch.DepartmentID, sch.LastExecutionDate, sch.NextExecutionDate, sch.IsActive,
	)
	return err
}

// FindByID returns schedule.
func (s *ScheduleStore) FindByID(ctx context.Context, id string) (*scheduleDomain.Schedule, error) {
	query := `SELECT id, template_id, frequency, inspector_id, department_id, last_execution_date, next_execution_date, is_active FROM inspection_schedule WHERE id = $1`
	sch := &scheduleDomain.Schedule{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&sch.ID, &sch.TemplateID, &sch.Frequency, &sch.InspectorID, &sch.DepartmentID, &sch.LastExecutionDate, &sch.NextExecutionDate, &sch.IsActive)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("schedule not found: %s", id)
		}
		return nil, err
	}
	return sch, nil
}

// ListActive retrieves enabled schedules.
func (s *ScheduleStore) ListActive(ctx context.Context) ([]*scheduleDomain.Schedule, error) {
	query := `SELECT id, template_id, frequency, inspector_id, department_id, last_execution_date, next_execution_date, is_active FROM inspection_schedule WHERE is_active = true`
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schedules []*scheduleDomain.Schedule
	for rows.Next() {
		sch := &scheduleDomain.Schedule{}
		err = rows.Scan(&sch.ID, &sch.TemplateID, &sch.Frequency, &sch.InspectorID, &sch.DepartmentID, &sch.LastExecutionDate, &sch.NextExecutionDate, &sch.IsActive)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, sch)
	}
	return schedules, nil
}

// Update saves recurrence dates.
func (s *ScheduleStore) Update(ctx context.Context, sch *scheduleDomain.Schedule) error {
	query := `UPDATE inspection_schedule SET last_execution_date = $2, next_execution_date = $3, is_active = $4 WHERE id = $1`
	_, err := s.db.ExecContext(ctx, query, sch.ID, sch.LastExecutionDate, sch.NextExecutionDate, sch.IsActive)
	return err
}
