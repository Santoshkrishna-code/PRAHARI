package middleware

import (
	"net/http"
)

// Chain holds a sequence of middleware handlers to run in FIFO order.
type Chain struct {
	middlewares []Middleware
}

// New constructs a new Chain.
func New(middlewares ...Middleware) Chain {
	return Chain{middlewares: append([]Middleware(nil), middlewares...)}
}

// Then chains the handler at the end of the pipeline, returning a single http.Handler.
func (c Chain) Then(handler http.Handler) http.Handler {
	if handler == nil {
		handler = http.DefaultServeMux
	}

	// Wrap in reverse order so the first middleware executes first
	for i := len(c.middlewares) - 1; i >= 0; i-- {
		handler = c.middlewares[i](handler)
	}

	return handler
}

// Append returns a new Chain with additional middlewares appended to the end.
func (c Chain) Append(middlewares ...Middleware) Chain {
	newM := make([]Middleware, 0, len(c.middlewares)+len(middlewares))
	newM = append(newM, c.middlewares...)
	newM = append(newM, middlewares...)
	return Chain{middlewares: newM}
}
