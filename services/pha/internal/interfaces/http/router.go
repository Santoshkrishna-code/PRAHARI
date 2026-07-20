package http

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /pha", h.CreatePHA)
	mux.HandleFunc("GET /pha", h.ListPHAs)
	mux.HandleFunc("GET /pha/{id}", h.GetPHA)
	mux.HandleFunc("POST /pha/{id}/hazop", h.ExecuteHAZOP)
	mux.HandleFunc("POST /pha/{id}/lopa", h.ExecuteLOPA)
	mux.HandleFunc("POST /pha/{id}/recommendation", h.CreateRecommendation)
	mux.HandleFunc("POST /pha/{id}/approval", h.ApprovePHA)
	mux.HandleFunc("POST /pha/search", h.SearchPHAs)
	mux.HandleFunc("GET /reports", h.GetExecutiveReport)
	mux.HandleFunc("GET /export/csv", h.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", h.ExportPDF)
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /ready", h.Health)
	mux.HandleFunc("GET /live", h.Health)
}
