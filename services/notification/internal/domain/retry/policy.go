package retry

// Policy holds execution counts configurations.
type Policy struct {
	MaxAttempts int `json:"max_attempts"`
	BackoffSec  int `json:"backoff_sec"`
}
