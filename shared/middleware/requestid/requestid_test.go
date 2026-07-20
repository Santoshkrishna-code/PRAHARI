package requestid_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	prahariMid "prahari/shared/middleware"
	"prahari/shared/middleware/requestid"
)

func TestRequestIDMiddleware_GeneratesNew(t *testing.T) {
	var capturedID string
	handler := requestid.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedID = prahariMid.GetCorrelationID(r.Context())
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	resp := rec.Result()
	responseHeaderID := resp.Header.Get(requestid.HeaderName)

	if capturedID == "" {
		t.Fatal("expected request context to contain generated correlation ID, got empty")
	}

	if responseHeaderID != capturedID {
		t.Errorf("expected response header %s to match context ID %s, got %s", requestid.HeaderName, capturedID, responseHeaderID)
	}
}

func TestRequestIDMiddleware_PropagatesExisting(t *testing.T) {
	var capturedID string
	handler := requestid.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedID = prahariMid.GetCorrelationID(r.Context())
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(requestid.HeaderName, "existing-id-xyz")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if capturedID != "existing-id-xyz" {
		t.Errorf("expected to propagate existing correlation ID, got %s", capturedID)
	}
}
