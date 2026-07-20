package logger

// Format defines output structure formats.
type Format string

const (
	FormatJSON    Format = "json"
	FormatConsole Format = "console"
)

// Options holds configuration settings for logger initialization.
type Options struct {
	Env         string   `json:"env"`          // e.g. development, production
	Level       string   `json:"level"`        // debug, info, warn, error, fatal
	Format      Format   `json:"format"`       // json, console
	OutputPaths []string `json:"output_paths"` // defaults to stdout
	ServiceName string   `json:"service_name"` // service name context
}

// DefaultOptions returns standard configuration for safety.
func DefaultOptions() Options {
	return Options{
		Env:         "development",
		Level:       "debug",
		Format:      FormatConsole,
		OutputPaths: []string{"stdout"},
		ServiceName: "prahari-service",
	}
}
