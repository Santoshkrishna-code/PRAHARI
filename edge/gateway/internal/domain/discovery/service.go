package discovery

// Service maps cluster-level service records names to downstream host arrays.
type Service struct {
	Name  string   `json:"name"`
	Hosts []string `json:"hosts"`
}
