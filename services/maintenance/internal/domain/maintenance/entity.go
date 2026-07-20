package maintenance

import (
	"errors"
	"time"
)

// Priority maps risk parameters.
type Priority string

const (
	PriorityLow       Priority = "LOW"
	PriorityMedium    Priority = "MEDIUM"
	PriorityHigh      Priority = "HIGH"
	PriorityEmergency Priority = "EMERGENCY"
)

// Maintenance is the central aggregate root of the Maintenance Management domain.
type Maintenance struct {
	ID                string     `json:"id" db:"id"`
	MaintenanceNumber string     `json:"maintenance_number" db:"maintenance_number"`
	AssetID           string     `json:"asset_id" db:"asset_id"`
	MaintenanceType   string     `json:"maintenance_type" db:"maintenance_type"`
	Priority          Priority   `json:"priority" db:"priority"`
	DepartmentID      string     `json:"department_id" db:"department_id"`
	Title             string     `json:"title" db:"title"`
	Description       string     `json:"description" db:"description"`
	StatusCode        string     `json:"status_code" db:"status_code"`
	TotalEstimatedCost float64   `json:"total_estimated_cost" db:"total_estimated_cost"`
	TotalActualCost   float64    `json:"total_actual_cost" db:"total_actual_cost"`
	CreatedAt         time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at" db:"updated_at"`
	IsDeleted         bool       `json:"is_deleted" db:"is_deleted"`
}

// Validate checks domain invariants for Maintenance.
func (m *Maintenance) Validate() error {
	if m.Title == "" {
		return errors.New("maintenance title is required")
	}
	if len(m.Title) > 200 {
		return errors.New("maintenance title must not exceed 200 characters")
	}
	if m.AssetID == "" {
		return errors.New("asset ID is required")
	}
	if m.DepartmentID == "" {
		return errors.New("department ID is required")
	}
	return nil
}
