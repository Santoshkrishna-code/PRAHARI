package object

// Object represents a tracked physical instance.
type Object struct {
	ID    string `json:"id"`
	Class string `json:"class"` // E.g., person, forklift
}
