package compatibility

// Rule checks if two chemical hazard groups can be stored in close proximity.
type Rule struct {
	ID             string `json:"id"`
	ClassA         string `json:"class_a"`
	ClassB         string `json:"class_b"`
	Compatible     bool   `json:"compatible"`
	SegregationReq string `json:"segregation_req,omitempty"` // E.g., Keep 3 meters apart
}
