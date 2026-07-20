package middleware

import (
	"net/http"
)

// Middleware is a standard decorator signature wrapping HTTP Handlers.
type Middleware func(http.Handler) http.Handler
