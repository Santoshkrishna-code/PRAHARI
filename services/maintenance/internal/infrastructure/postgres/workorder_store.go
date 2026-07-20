package postgres

import (
	"context"
	"database/sql"
	"fmt"

	workorderDomain "prahari/services/maintenance/internal/domain/workorder"
)

// WorkOrderStore implements work order persistence operations.
type WorkOrderStore struct {
	db *sql.DB
}

// NewWorkOrderStore instantiates WorkOrderStore.
func NewWorkOrderStore(db *sql.DB) *WorkOrderStore {
	return &WorkOrderStore{db: db}
}

// Create persists a work order.
func (s *WorkOrderStore) Create(ctx context.Context, w *workorderDomain.WorkOrder) error {
	query := `INSERT INTO work_orders (id, maintenance_id, work_order_number, scheduled_date, estimated_hours, actual_hours)
		VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.ExecContext(ctx, query, w.ID, w.MaintenanceID, w.WorkOrderNumber, w.ScheduledDate, w.EstimatedHours, w.ActualHours)
	return err
}

// FindByID returns a work order.
func (s *WorkOrderStore) FindByID(ctx context.Context, id string) (*workorderDomain.WorkOrder, error) {
	query := `SELECT id, maintenance_id, work_order_number, scheduled_date, actual_start_date, actual_end_date, completed_by, estimated_hours, actual_hours FROM work_orders WHERE id = $1`
	w := &workorderDomain.WorkOrder{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&w.ID, &w.MaintenanceID, &w.WorkOrderNumber, &w.ScheduledDate, &w.ActualStartDate, &w.ActualEndDate, &w.CompletedBy, &w.EstimatedHours, &w.ActualHours)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("work order not found: %s", id)
		}
		return nil, err
	}
	return w, nil
}

// FindByMaintenanceID returns work orders generated for a maintenance profile.
func (s *WorkOrderStore) FindByMaintenanceID(ctx context.Context, maintenanceID string) ([]*workorderDomain.WorkOrder, error) {
	query := `SELECT id, maintenance_id, work_order_number, scheduled_date, actual_start_date, actual_end_date, completed_by, estimated_hours, actual_hours FROM work_orders WHERE maintenance_id = $1`
	rows, err := s.db.QueryContext(ctx, query, maintenanceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []*workorderDomain.WorkOrder
	for rows.Next() {
		w := &workorderDomain.WorkOrder{}
		err = rows.Scan(&w.ID, &w.MaintenanceID, &w.WorkOrderNumber, &w.ScheduledDate, &w.ActualStartDate, &w.ActualEndDate, &w.CompletedBy, &w.EstimatedHours, &w.ActualHours)
		if err != nil {
			return nil, err
		}
		list = append(list, w)
	}
	return list, nil
}
