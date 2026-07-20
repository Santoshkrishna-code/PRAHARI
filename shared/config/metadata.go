package config

// Metadata stores runtime telemetry build attributes for diagnostic checks.
type Metadata struct {
	Version     string `json:"version"`
	BuildTime   string `json:"build_time"`
	GitCommit   string `json:"git_commit"`
	ServiceName string `json:"service_name"`
}

var (
	// Global build metadata populated during compilations using -ldflags
	Version   = "1.0.0"
	BuildTime = "unknown"
	GitCommit = "unknown"
)

// GetMetadata returns compiled service version tags.
func GetMetadata(serviceName string) Metadata {
	return Metadata{
		Version:     Version,
		BuildTime:   BuildTime,
		GitCommit:   GitCommit,
		ServiceName: serviceName,
	}
}
