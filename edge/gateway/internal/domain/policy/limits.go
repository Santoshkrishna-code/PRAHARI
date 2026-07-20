package policy

// Policy defines parameters mapping request limits counts.
type Policy struct {
	RequestsPerMin int64 `json:"requests_per_min"`
}
