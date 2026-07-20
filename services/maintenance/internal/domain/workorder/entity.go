package workorder

import (
	"errors"
	"time"
)

// WorkOrder represents the actual execution card details for technicians.
type WorkOrder struct {
	ID              string     `json:"id" db:"id"`
	MaintenanceID   string     `json:"maintenance_id" db:"maintenance_id"`
	WorkOrderNumber string     `json:"work_order_number" db:"work_order_number"`
	ScheduledDate   time.Time  `json:"scheduled_date" db:"scheduled_date"`
	ActualStartDate *time.Time `json:"actual_start_date,omitempty" db:"actual_start_date"`
	ActualEndDate   *time.Time `json:"actual_end_date,omitempty" db:"actual_end_date"`
	CompletedBy     string     `json:"completed_by,omitempty" db:"completed_by"`
	EstimatedHours  float64    `json:"estimated_hours" db:"estimated_hours"`
	ActualHours     float64    `json:"actual_hours" db:"actual_hours"`
}

// Validate checks domain invariants.
func (w *WorkOrder) Validate() error {
	if w.MaintenanceID == "" {
		return errors.New("maintenance ID reference is required")
	}
	if w.WorkOrderNumber == "" {
		return errors.New("work order number code is required")
	}
	return nil
}
