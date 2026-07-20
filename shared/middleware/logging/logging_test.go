package logging_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"prahari/shared/middleware/logging"
)

func TestLoggingMiddleware(t *testing.T) {
	handler := logging.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
	}))

	req := httptest.NewRequest(http.MethodPost, "/submit", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusAccepted {
		t.Errorf("expected 202 Accepted, got %d", rec.Code)
	}
}
