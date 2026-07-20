package resilience

import (
	"time"
)

// CBConfig holds configurations for circuit breaker instances.
type CBConfig struct {
	FailureRatio     float64       `json:"failure_ratio"` // e.g. 0.5 (50% failure rate)
	CooldownDuration time.Duration `json:"cooldown_duration"`
}

// Config maps default parameters for execution wrappers.
type Config struct {
	Timeout       time.Duration `json:"timeout"`
	BulkheadLimit int           `json:"bulkhead_limit"` // Semaphore concurrency limit
	CB            CBConfig      `json:"cb"`
}
