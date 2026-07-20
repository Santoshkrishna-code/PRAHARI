package http

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /chemicals/receive", h.ReceiveChemical)
	mux.HandleFunc("POST /chemicals/{id}/issue", h.IssueChemical)
	mux.HandleFunc("POST /chemicals/{id}/return", h.ReturnChemical)
	mux.HandleFunc("POST /chemicals/{id}/transfer", h.TransferChemical)
	mux.HandleFunc("GET /chemicals/{id}/sds", h.GetSDS)
	mux.HandleFunc("POST /chemicals/sds", h.RegisterSDS)
	mux.HandleFunc("GET /chemicals", h.ListChemicals)
	mux.HandleFunc("GET /chemicals/{id}", h.GetChemical)
	mux.HandleFunc("POST /chemicals/search", h.SearchChemicals)
	mux.HandleFunc("GET /reports", h.GetExecutiveReport)
	mux.HandleFunc("GET /export/csv", h.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", h.ExportPDF)
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /ready", h.Health)
	mux.HandleFunc("GET /live", h.Health)
}
