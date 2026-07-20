package idempotency

import (
	"net/http"
)

const (
	HeaderName = "Idempotency-Key"
)

// GetIdempotencyKey extracts the custom Idempotency-Key header value from HTTP requests.
func GetIdempotencyKey(r *http.Request) string {
	return r.Header.Get(HeaderName)
}
