package http

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("GET /dashboards", h.GetDashboard)
	mux.HandleFunc("POST /dashboards", h.CreateDashboard)
	mux.HandleFunc("GET /kpis", h.GetKPIs)
	mux.HandleFunc("GET /metrics", h.GetMetrics)
	mux.HandleFunc("POST /reports", h.CreateReport)
	mux.HandleFunc("GET /reports/{id}", h.GetReport)
	mux.HandleFunc("POST /analytics/search", h.SearchAnalytics)
	mux.HandleFunc("GET /export/csv", h.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", h.ExportPDF)
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /ready", h.Health)
	mux.HandleFunc("GET /live", h.Health)
}
