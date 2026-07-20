package builders_test

import (
	"net/http"
	"testing"

	"prahari/shared/testing/builders"
)

func TestHTTPBuilders(t *testing.T) {
	// 1. BuildRequest
	req1 := builders.BuildRequest(http.MethodGet, "/list", []byte(`{"id":1}`))
	if req1.Method != http.MethodGet || req1.URL.Path != "/list" {
		t.Errorf("mismatched request parameters: %+v", req1)
	}

	// 2. BuildJSONRequest
	payload := map[string]string{"name": "prahari"}
	req2 := builders.BuildJSONRequest(http.MethodPost, "/add", payload)
	if req2.Method != http.MethodPost || req2.Header.Get("Content-Type") != "application/json" {
		t.Errorf("mismatched JSON request parameters: %+v", req2)
	}

	// 3. ExecuteRequest
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusAccepted)
	})
	rec := builders.ExecuteRequest(h, req1)
	if rec.Code != http.StatusAccepted {
		t.Errorf("expected 202 Accepted response, got %d", rec.Code)
	}
}
