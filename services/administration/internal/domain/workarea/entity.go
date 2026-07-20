package workarea

import "time"

// WorkArea represents a designated workspace area (e.g. welding booth, chemical line).
type WorkArea struct {
	ID           string    `json:"id"`
	DepartmentID string    `json:"department_id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	HazardLevel  string    `json:"hazard_level"` // LOW, MEDIUM, HIGH
	CreatedAt    time.Time `json:"created_at"`
}
