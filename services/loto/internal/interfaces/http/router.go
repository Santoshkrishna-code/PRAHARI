package http

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /loto", h.CreateLOTO)
	mux.HandleFunc("GET /loto", h.ListLOTOs)
	mux.HandleFunc("GET /loto/{id}", h.GetLOTO)
	mux.HandleFunc("POST /loto/{id}/approve", h.ApproveIsolation)
	mux.HandleFunc("POST /loto/{id}/apply-locks", h.ApplyLocks)
	mux.HandleFunc("POST /loto/{id}/verify-zero-energy", h.VerifyZeroEnergy)
	mux.HandleFunc("POST /loto/{id}/restore", h.RestoreSystem)
	mux.HandleFunc("POST /loto/search", h.SearchLOTOs)
	mux.HandleFunc("GET /reports", h.GetExecutiveReport)
	mux.HandleFunc("GET /export/csv", h.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", h.ExportPDF)
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /ready", h.Health)
	mux.HandleFunc("GET /live", h.Health)
}
