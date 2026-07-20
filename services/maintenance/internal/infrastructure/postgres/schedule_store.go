package postgres

import (
	"context"
	"database/sql"
	"fmt"

	scheduleDomain "prahari/services/maintenance/internal/domain/schedule"
)

// ScheduleStore implements schedules calendar entries persistence.
type ScheduleStore struct {
	db *sql.DB
}

// NewScheduleStore instantiates ScheduleStore.
func NewScheduleStore(db *sql.DB) *ScheduleStore {
	return &ScheduleStore{db: db}
}

// Create persists schedule settings.
func (s *ScheduleStore) Create(ctx context.Context, sch *scheduleDomain.Schedule) error {
	query := `INSERT INTO maintenance_schedules (id, maintenance_id, scheduled_start_date, scheduled_end_date, estimated_downtime_min)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, sch.ID, sch.MaintenanceID, sch.ScheduledStartDate, sch.ScheduledEndDate, sch.EstimatedDowntimeMin)
	return err
}

// FindByID returns schedule.
func (s *ScheduleStore) FindByID(ctx context.Context, id string) (*scheduleDomain.Schedule, error) {
	query := `SELECT id, maintenance_id, scheduled_start_date, scheduled_end_date, estimated_downtime_min FROM maintenance_schedules WHERE id = $1`
	sch := &scheduleDomain.Schedule{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&sch.ID, &sch.MaintenanceID, &sch.ScheduledStartDate, &sch.ScheduledEndDate, &sch.EstimatedDowntimeMin)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("schedule not found: %s", id)
		}
		return nil, err
	}
	return sch, nil
}
