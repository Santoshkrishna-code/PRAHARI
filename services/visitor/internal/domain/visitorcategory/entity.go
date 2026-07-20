package visitorcategory

import "time"

// Category represents structured categories of plant visitors.
type Category struct {
	ID          string    `json:"id"`
	Code        string    `json:"code"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	InductionReq bool     `json:"induction_required"`
	CreatedAt   time.Time `json:"created_at"`
}
