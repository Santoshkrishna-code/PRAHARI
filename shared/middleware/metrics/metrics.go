package metrics

import (
	"net/http"
	"sync/atomic"
	"time"
)

// Tracker aggregates metrics counters.
type Tracker struct {
	RequestCount int64
	LatencySumNs int64
}

// NewTracker constructs a metrics Tracker.
func NewTracker() *Tracker {
	return &Tracker{}
}

// Middleware records processing statistics for each incoming request.
func (t *Tracker) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		atomic.AddInt64(&t.RequestCount, 1)
		atomic.AddInt64(&t.LatencySumNs, time.Since(start).Nanoseconds())
	})
}
