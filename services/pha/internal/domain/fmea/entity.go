package fmea

import "time"

// Analysis represents a Failure Modes and Effects Analysis.
type Analysis struct {
	ID             string    `json:"id"`
	StudyID        string    `json:"study_id"`
	EquipmentItem  string    `json:"equipment_item"`
	FailureMode    string    `json:"failure_mode"`
	FailureEffect  string    `json:"failure_effect"`
	Severity       int       `json:"severity"`       // 1 to 10
	Occurrence     int       `json:"occurrence"`     // 1 to 10
	Detection      int       `json:"detection"`      // 1 to 10
	RPN            int       `json:"rpn"`            // Risk Priority Number = Severity * Occurrence * Detection
	CreatedAt      time.Time `json:"created_at"`
}
