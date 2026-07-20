package phastudy

import "time"

// MethodType represents the analysis technique.
type MethodType string

const (
	MethodHAZOP     MethodType = "HAZOP"
	MethodLOPA      MethodType = "LOPA"
	MethodBowTie    MethodType = "BOWTIE"
	MethodFMEA      MethodType = "FMEA"
	MethodWhatIf    MethodType = "WHAT_IF"
	MethodChecklist MethodType = "CHECKLIST"
	MethodPSSR      MethodType = "PSSR"
)

// Study represents a formal Process Hazard Analysis (PHA) study.
type Study struct {
	ID                 string     `json:"id"`
	StudyNumber        string     `json:"study_number"`
	PlantID            string     `json:"plant_id"`
	UnitID             string     `json:"unit_id"`
	Title              string     `json:"title"`
	Description        string     `json:"description"`
	Method             MethodType `json:"method"`
	MOCID              string     `json:"moc_id,omitempty"` // Triggering MOC if applicable
	Status             string     `json:"status"`           // Draft, Preparation, Study, Risk Evaluation, Recommendation, Approval, Implementation Tracking, Verification, Revalidation Scheduled, Closed, Cancelled, Superseded
	LeaderID           string     `json:"leader_id"`
	ScribeID           string     `json:"scribe_id"`
	TargetDate         time.Time  `json:"target_date"`
	RevalidationDueAt *time.Time `json:"revalidation_due_at,omitempty"` // Typically 5 years per OSHA PSM
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
}
