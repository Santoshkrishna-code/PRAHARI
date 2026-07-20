package http

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /ppe", h.CreatePPE)
	mux.HandleFunc("GET /ppe", h.ListPPEs)
	mux.HandleFunc("GET /ppe/{id}", h.GetPPE)
	mux.HandleFunc("POST /ppe/{id}/issue", h.IssuePPE)
	mux.HandleFunc("POST /ppe/{id}/return", h.ReturnPPE)
	mux.HandleFunc("POST /ppe/{id}/inspect", h.InspectPPE)
	mux.HandleFunc("POST /ppe/{id}/maintain", h.MaintainPPE)
	mux.HandleFunc("POST /ppe/search", h.SearchPPEs)
	mux.HandleFunc("GET /reports", h.GetExecutiveReport)
	mux.HandleFunc("GET /export/csv", h.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", h.ExportPDF)
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /ready", h.Health)
	mux.HandleFunc("GET /live", h.Health)
}
