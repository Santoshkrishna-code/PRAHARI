package builders

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

// BuildRequest instantiates a generic request with byte content.
func BuildRequest(method, path string, body []byte) *http.Request {
	var bodyReader *bytes.Reader
	if body != nil {
		bodyReader = bytes.NewReader(body)
	}

	req := httptest.NewRequest(method, path, bodyReader)
	req.Header.Set("Content-Type", "application/json")
	return req
}

// BuildJSONRequest parses structure maps into JSON bytes, return request object.
func BuildJSONRequest(method, path string, payload interface{}) *http.Request {
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		bodyBytes = nil
	}
	return BuildRequest(method, path, bodyBytes)
}

// ExecuteRequest runs the target HTTP handler, returning the response recorder object.
func ExecuteRequest(handler http.Handler, req *http.Request) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	return rec
}
