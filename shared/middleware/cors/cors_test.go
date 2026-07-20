package cors_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"prahari/shared/middleware/cors"
)

func TestCORSMiddleware_Preflight(t *testing.T) {
	opts := cors.DefaultOptions()
	middleware := cors.Middleware(opts)

	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodOptions, "/items", nil)
	req.Header.Set("Origin", "https://prahari.gov")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	resp := rec.Result()
	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("expected 204 No Content for preflight OPTIONS, got %d", resp.StatusCode)
	}

	allowOrigin := resp.Header.Get("Access-Control-Allow-Origin")
	if allowOrigin != "https://prahari.gov" {
		t.Errorf("expected Access-Control-Allow-Origin to be 'https://prahari.gov', got '%s'", allowOrigin)
	}
}
