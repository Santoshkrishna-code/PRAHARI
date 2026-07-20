package http

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /actions", h.CreateAction)
	mux.HandleFunc("GET /actions", h.ListActions)
	mux.HandleFunc("GET /actions/{id}", h.GetAction)
	mux.HandleFunc("POST /actions/{id}/assign", h.AssignAction)
	mux.HandleFunc("POST /actions/{id}/submit-evidence", h.SubmitEvidence)
	mux.HandleFunc("POST /actions/{id}/review", h.ReviewAction)
	mux.HandleFunc("POST /actions/{id}/close", h.CloseAction)
	mux.HandleFunc("POST /actions/search", h.SearchActions)
	mux.HandleFunc("GET /reports", h.GetExecutiveReport)
	mux.HandleFunc("GET /export/csv", h.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", h.ExportPDF)
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /ready", h.Health)
	mux.HandleFunc("GET /live", h.Health)
}
