package sources

import (
	"github.com/caarlos0/env/v9"
)

// ParseEnv parses standard environment variable values directly into the target configuration struct.
func ParseEnv(target interface{}) error {
	opts := env.Options{
		RequiredIfNoDef: true, // Fail-fast security check if required is missing
	}
	return env.ParseWithOptions(target, opts)
}
