package pssr

import "time"

// Review represents a Pre-Startup Safety Review (PSSR) verification item.
type Review struct {
	ID              string    `json:"id"`
	StudyID         string    `json:"study_id"`
	MOCID           string    `json:"moc_id,omitempty"`
	PSSRTitle       string    `json:"pssr_title"`
	ConstructionOK bool      `json:"construction_ok"`
	ProceduresOK   bool      `json:"procedures_ok"`
	TrainingOK     bool      `json:"training_ok"`
	PHAActionItemsOK bool    `json:"pha_action_items_ok"`
	Status          string    `json:"status"` // APPROVED, REJECTED, PENDING
	VerifiedBy      string    `json:"verified_by"`
	VerifiedAt      time.Time `json:"verified_at"`
}
