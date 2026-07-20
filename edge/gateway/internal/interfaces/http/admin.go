package http

import (
	"encoding/json"
	"net/http"

	"prahari/edge/gateway/internal/application/routing"
)

// AdminHandler maps routes configuration controllers.
type AdminHandler struct {
	table *routing.Table
}

// NewAdminHandler constructs an AdminHandler.
func NewAdminHandler(t *routing.Table) *AdminHandler {
	return &AdminHandler{table: t}
}

// ListRoutes handles GET requests listing active gateway routes.
func (h *AdminHandler) ListRoutes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status": "gateway is online",
	})
}
