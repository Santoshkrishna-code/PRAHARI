package http

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /barriers", h.CreateBarrier)
	mux.HandleFunc("GET /barriers", h.ListBarriers)
	mux.HandleFunc("GET /barriers/{id}", h.GetBarrier)
	mux.HandleFunc("POST /barriers/{id}/proof-test", h.RecordProofTest)
	mux.HandleFunc("POST /barriers/{id}/integrity", h.AssessIntegrity)
	mux.HandleFunc("POST /barriers/{id}/impairment", h.RegisterImpairment)
	mux.HandleFunc("POST /barriers/{id}/bypass", h.RegisterBypass)
	mux.HandleFunc("POST /barriers/search", h.SearchBarriers)
	mux.HandleFunc("GET /reports", h.GetExecutiveReport)
	mux.HandleFunc("GET /export/csv", h.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", h.ExportPDF)
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /ready", h.Health)
	mux.HandleFunc("GET /live", h.Health)
}
