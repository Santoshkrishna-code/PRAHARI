package upstream

// Host represents a target downstream server instance.
type Host struct {
	Address string `json:"address"`
	Healthy bool   `json:"healthy"`
}
