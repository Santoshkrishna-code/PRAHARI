package deployment

import "time"

// Configuration represents model version deployment metrics to Edge hardware.
type Configuration struct {
	ID         string    `json:"id"`
	EdgeDevice string    `json:"edge_device"` // E.g., Jetson-Nano-02
	ModelID    string    `json:"model_id"`
	Version    string    `json:"version"`
	DeployedAt time.Time `json:"deployed_at"`
}
