package policy

import (
	"math"
	"time"

	"prahari/services/integration/internal/domain/retry"
)

// CalculateNextBackoff calculates backoff delay using exponential backoff strategy.
func CalculateNextBackoff(retryCount int) time.Duration {
	// E.g. 2^retryCount * 5 seconds. Max delay capped at 1 hour.
	backoffSec := math.Pow(2, float64(retryCount)) * 5
	if backoffSec > 3600 {
		backoffSec = 3600
	}
	return time.Duration(backoffSec) * time.Second
}

// IsRetryExhausted checks if retry count has exceeded maximum retries.
func IsRetryExhausted(m *retry.Message) bool {
	if m == nil {
		return true
	}
	return m.RetryCount >= m.MaxRetries
}
