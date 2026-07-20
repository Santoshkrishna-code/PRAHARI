package rootcause

import "time"

// Analysis represents a systematic investigation of a safety failure or deviation (supports 5-Why and Fishbone classifications).
type Analysis struct {
	ID            string    `json:"id"`
	CapaID        string    `json:"capa_id"`
	Method        string    `json:"method"` // FIVE_WHYS, FISHBONE, FAULT_TREE
	FindingsText  string    `json:"findings_text"`
	Why1          string    `json:"why_1,omitempty"`
	Why2          string    `json:"why_2,omitempty"`
	Why3          string    `json:"why_3,omitempty"`
	Why4          string    `json:"why_4,omitempty"`
	Why5          string    `json:"why_5,omitempty"`
	RootCauseType string    `json:"root_cause_type"` // HUMAN_ERROR, EQUIPMENT_FAILURE, PROCESS_DEFICIENCY
	AnalyzedBy    string    `json:"analyzed_by"`
	AnalyzedAt    time.Time `json:"analyzed_at"`
}
