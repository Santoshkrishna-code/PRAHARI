package metrics_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"prahari/shared/middleware/metrics"
)

func TestMetricsMiddleware(t *testing.T) {
	tracker := metrics.NewTracker()
	middleware := tracker.Middleware

	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if tracker.RequestCount != 1 {
		t.Errorf("expected request count to be 1, got %d", tracker.RequestCount)
	}

	if tracker.LatencySumNs <= 0 {
		t.Errorf("expected positive latency logs, got %d", tracker.LatencySumNs)
	}
}
