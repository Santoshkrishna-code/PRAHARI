package recovery_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"prahari/shared/middleware/recovery"
)

func TestRecoveryMiddleware_InterceptsPanics(t *testing.T) {
	handler := recovery.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("database execution failure")
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	// Should not crash the test suite
	handler.ServeHTTP(rec, req)

	resp := rec.Result()
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected 500 Internal Server Error, got %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "application/problem+json" {
		t.Errorf("expected application/problem+json content type, got '%s'", contentType)
	}
}
