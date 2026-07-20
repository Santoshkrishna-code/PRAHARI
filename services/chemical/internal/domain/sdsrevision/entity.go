package sdsrevision

import "time"

// Revision logs history changes of Safety Data Sheets.
type Revision struct {
	ID          string    `json:"id"`
	SdsID       string    `json:"sds_id"`
	RevisionNum string    `json:"revision_num"`
	RevisedBy   string    `json:"revised_by"`
	RevisedAt   time.Time `json:"revised_at"`
	ChangeLog   string    `json:"change_log"`
	DocumentURL string    `json:"document_url"`
}
