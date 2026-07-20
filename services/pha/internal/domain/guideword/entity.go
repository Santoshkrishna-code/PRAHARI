package guideword

import "time"

// Word represents a standard HAZOP guide word definition.
type Word struct {
	ID          string    `json:"id"`
	Word        string    `json:"word"` // HIGH, LOW, NO, REVERSE, OTHER THAN, AS WELL AS
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
