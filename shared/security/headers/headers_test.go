package headers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"prahari/shared/security/headers"
)

func TestSecurityHeadersMiddleware(t *testing.T) {
	opts := headers.DefaultOptions()
	middleware := headers.Middleware(opts)

	finalHandler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	finalHandler.ServeHTTP(rec, req)

	resp := rec.Result()

	// Assertions for standard security headers
	if xfo := resp.Header.Get("X-Frame-Options"); xfo != "DENY" {
		t.Errorf("expected X-Frame-Options DENY, got '%s'", xfo)
	}

	if xct := resp.Header.Get("X-Content-Type-Options"); xct != "nosniff" {
		t.Errorf("expected X-Content-Type-Options nosniff, got '%s'", xct)
	}

	if hsts := resp.Header.Get("Strict-Transport-Security"); hsts == "" {
		t.Error("expected Strict-Transport-Security header to be injected")
	}

	if csp := resp.Header.Get("Content-Security-Policy"); csp == "" {
		t.Error("expected Content-Security-Policy header to be injected")
	}

	if perm := resp.Header.Get("Permissions-Policy"); perm == "" {
		t.Error("expected Permissions-Policy header to be injected")
	}

	if coep := resp.Header.Get("Cross-Origin-Embedder-Policy"); coep != "require-corp" {
		t.Errorf("expected Cross-Origin-Embedder-Policy require-corp, got '%s'", coep)
	}
}
