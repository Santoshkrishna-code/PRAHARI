package http

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /connectors", h.RegisterConnector)
	mux.HandleFunc("GET /connectors", h.ListConnectors)
	mux.HandleFunc("POST /connectors/test", h.TestConnector)
	mux.HandleFunc("POST /jobs", h.RunSyncJob)
	mux.HandleFunc("POST /jobs/{id}/run", h.RunSyncJob)
	mux.HandleFunc("POST /webhooks", h.CreateWebhook)
	mux.HandleFunc("POST /integration/search", h.SearchIntegrations)
	mux.HandleFunc("GET /reports", h.GetExecutiveReport)
	mux.HandleFunc("GET /export/csv", h.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", h.ExportPDF)
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /ready", h.Health)
	mux.HandleFunc("GET /live", h.Health)
}
