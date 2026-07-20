package postgres

import (
	"context"
	"database/sql"

	taskDomain "prahari/services/maintenance/internal/domain/task"
)

// TaskStore implements individual work steps checklist query ports.
type TaskStore struct {
	db *sql.DB
}

// NewTaskStore instantiates TaskStore.
func NewTaskStore(db *sql.DB) *TaskStore {
	return &TaskStore{db: db}
}

// Create inserts task.
func (s *TaskStore) Create(ctx context.Context, t *taskDomain.Task) error {
	query := `INSERT INTO maintenance_tasks (id, maintenance_id, description, sequence_order, is_completed)
		VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.ExecContext(ctx, query, t.ID, t.MaintenanceID, t.Description, t.SequenceOrder, t.IsCompleted)
	return err
}

// FindByMaintenanceID returns checklist elements list.
func (s *TaskStore) FindByMaintenanceID(ctx context.Context, maintenanceID string) ([]*taskDomain.Task, error) {
	query := `SELECT id, maintenance_id, description, sequence_order, is_completed FROM maintenance_tasks WHERE maintenance_id = $1 ORDER BY sequence_order ASC`
	rows, err := s.db.QueryContext(ctx, query, maintenanceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*taskDomain.Task
	for rows.Next() {
		t := &taskDomain.Task{}
		err = rows.Scan(&t.ID, &t.MaintenanceID, &t.Description, &t.SequenceOrder, &t.IsCompleted)
		if err != nil {
			return nil, err
		}
		list = append(list, t)
	}
	return list, nil
}
