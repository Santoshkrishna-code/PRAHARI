package http

import (
	"net/http"

	prahariErrors "prahari/shared/errors"
	"prahari/edge/gateway/internal/application/proxy"
	"prahari/edge/gateway/internal/application/routing"
)

// Handler maps the entry request controller matching path configurations.
type Handler struct {
	table *routing.Table
}

// NewHandler constructs a Handler.
func NewHandler(t *routing.Table) *Handler {
	return &Handler{table: t}
}

// ServeHTTP acts as the reverse proxy router interface.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Match requests path prefix
	rt, err := h.table.Match(r.URL.Path)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewNotFoundError("target path route not found", err))
		return
	}

	// Forward payload down the HTTP reverse proxy
	err = proxy.Forward(w, r, rt)
	if err != nil {
		prahariErrors.WriteHTTP(w, prahariErrors.NewInternalError("failed to proxy payload requests", err))
		return
	}
}
