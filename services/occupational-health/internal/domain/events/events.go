package events

import (
	"time"
)

// HealthProfileCreated represents profile registration events.
type HealthProfileCreated struct {
	ID         string    `json:"id"`
	WorkerID   string    `json:"worker_id"`
	WorkerType string    `json:"worker_type"`
	Timestamp  time.Time `json:"timestamp"`
}

// MedicalExaminationCompleted represents completion of health examinations.
type MedicalExaminationCompleted struct {
	ID              string    `json:"id"`
	HealthProfileID string    `json:"health_profile_id"`
	ExamType        string    `json:"exam_type"`
	OutcomeStatus   string    `json:"outcome_status"`
	Timestamp       time.Time `json:"timestamp"`
}

// FitnessAssessed represents assessment triggers.
type FitnessAssessed struct {
	ID              string    `json:"id"`
	HealthProfileID string    `json:"health_profile_id"`
	ResultCode      string    `json:"result_code"`
	Timestamp       time.Time `json:"timestamp"`
}

// MedicalClearanceGranted represents clearance events.
type MedicalClearanceGranted struct {
	ID              string    `json:"id"`
	HealthProfileID string    `json:"health_profile_id"`
	ExpiryDate      time.Time `json:"expiry_date"`
	Timestamp       time.Time `json:"timestamp"`
}

// MedicalRestrictionApplied represents application of medical restrictions.
type MedicalRestrictionApplied struct {
	ID              string    `json:"id"`
	HealthProfileID string    `json:"health_profile_id"`
	RestrictionCode string    `json:"restriction_code"`
	Timestamp       time.Time `json:"timestamp"`
}
