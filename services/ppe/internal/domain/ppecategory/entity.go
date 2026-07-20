package ppecategory

import "time"

// Category represents high-level classification of protective gear (e.g. Head, Eye, Fall Protection).
type Category struct {
	ID          string    `json:"id"`
	Code        string    `json:"code"` // E.g. HEAD, EYE, FALL, RESPIRATORY
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
