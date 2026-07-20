package search

import (
	"time"
)

// Criteria structures search params.
type Criteria struct {
	WorkerID          string     `json:"worker_id"`
	WorkerType        string     `json:"worker_type"`
	DepartmentID      string     `json:"department_id"`
	MedicalStatus     string     `json:"medical_status"`
	ClearanceStatus   string     `json:"clearance_status"`
	RestrictionCode   string     `json:"restriction_code"`
	VaccineName       string     `json:"vaccine_name"`
	IllnessName       string     `json:"illness_name"`
	PhysicianID       string     `json:"physician_id"`
	ExaminationType   string     `json:"examination_type"`
	StartDate         *time.Time `json:"start_date"`
	EndDate           *time.Time `json:"end_date"`
	FreeText          string     `json:"free_text"`
	Limit             int        `json:"limit"`
	Offset            int        `json:"offset"`
}
