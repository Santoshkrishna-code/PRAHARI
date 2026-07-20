package timeout

import (
	"net/http"
	"time"
)

// Middleware wraps the handler with a duration timeout limit.
// Returns a 503 Service Unavailable response if execution times out.
func Middleware(timeout time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.TimeoutHandler(
			next,
			timeout,
			`{"type":"about:blank","title":"Service Unavailable","status":503,"detail":"The request execution time exceeded limit boundaries."}`,
		)
	}
}
