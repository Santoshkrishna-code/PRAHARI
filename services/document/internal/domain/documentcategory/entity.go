package documentcategory

import "time"

// Category classifies documents (e.g., EHS SOPs, Operations, Mechanical Integrity, Quality).
type Category struct {
	ID          string    `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ParentID    string    `json:"parent_id,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}
