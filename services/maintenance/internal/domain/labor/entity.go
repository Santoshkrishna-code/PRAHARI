package labor

import (
	"errors"
)

// Labor tracks technician billings hours.
type Labor struct {
	ID            string  `json:"id" db:"id"`
	MaintenanceID string  `json:"maintenance_id" db:"maintenance_id"`
	TechnicianID  string  `json:"technician_id" db:"technician_id"`
	HoursWorked   float64 `json:"hours_worked" db:"hours_worked"`
	HourlyRate    float64 `json:"hourly_rate" db:"hourly_rate"`
}

// Validate checks domain invariants.
func (l *Labor) Validate() error {
	if l.MaintenanceID == "" {
		return errors.New("maintenance ID is required")
	}
	if l.TechnicianID == "" {
		return errors.New("technician ID is required")
	}
	if l.HoursWorked <= 0.0 {
		return errors.New("hours worked must be greater than zero")
	}
	return nil
}

// Cost calculated labor value sums.
func (l *Labor) Cost() float64 {
	return l.HoursWorked * l.HourlyRate
}
