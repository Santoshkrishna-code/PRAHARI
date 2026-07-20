package retry

import (
	"math"
	"math/rand"
	"time"
)

// Backoff configures exponential delays backoff policies.
type Backoff struct {
	Min    time.Duration
	Max    time.Duration
	Factor float64
	Jitter bool
}

// Duration computes backoff duration matching the attempt count, adding random jitter.
func (b Backoff) Duration(attempt int) time.Duration {
	if b.Factor <= 0 {
		b.Factor = 2.0
	}

	tempVal := float64(b.Min) * math.Pow(b.Factor, float64(attempt))
	dur := time.Duration(tempVal)

	if dur > b.Max {
		dur = b.Max
	}

	if b.Jitter {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		dur = time.Duration(r.Int63n(int64(dur)))
	}

	return dur
}
