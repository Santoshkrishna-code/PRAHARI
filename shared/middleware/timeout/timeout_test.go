package timeout_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"prahari/shared/middleware/timeout"
)

func TestTimeoutMiddleware_Triggers(t *testing.T) {
	middleware := timeout.Middleware(10 * time.Millisecond)
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(50 * time.Millisecond) // Exceeds timeout
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	resp := rec.Result()
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Errorf("expected 503 Service Unavailable, got %d", resp.StatusCode)
	}
}
