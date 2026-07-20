package attachment

import "time"

// Attachment represents a file linked to a water management record.
type Attachment struct {
	ID         string    `json:"id"`
	TargetType string    `json:"target_type"`
	TargetID   string    `json:"target_id"`
	FileName   string    `json:"file_name"`
	FileURL    string    `json:"file_url"`
	UploadedBy string    `json:"uploaded_by"`
	UploadedAt time.Time `json:"uploaded_at"`
}
