package checklist

import "time"

// Analysis represents a checklist-based process safety compliance study.
type Analysis struct {
	ID             string    `json:"id"`
	StudyID        string    `json:"study_id"`
	ChecklistGroup string    `json:"checklist_group"` // E.g., Pressure Vessels, Relief Systems
	ItemQuestion   string    `json:"item_question"`
	IsCompliant    bool      `json:"is_compliant"`
	Comments       string    `json:"comments"`
	CreatedAt      time.Time `json:"created_at"`
}
