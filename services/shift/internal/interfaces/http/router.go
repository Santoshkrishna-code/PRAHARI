package http

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /shifts", h.CreateShift)
	mux.HandleFunc("GET /shifts", h.ListShifts)
	mux.HandleFunc("GET /shifts/{id}", h.GetShift)
	mux.HandleFunc("POST /shifts/{id}/start", h.StartShift)
	mux.HandleFunc("POST /shifts/{id}/handover", h.InitiateHandover)
	mux.HandleFunc("POST /shifts/{id}/accept", h.AcceptHandover)
	mux.HandleFunc("POST /shifts/{id}/close", h.CloseShift)
	mux.HandleFunc("POST /shifts/{id}/log", h.LogActivity)
	mux.HandleFunc("POST /shifts/{id}/journal", h.WriteJournal)
	mux.HandleFunc("POST /shifts/search", h.SearchShifts)
	mux.HandleFunc("GET /reports", h.GetExecutiveReport)
	mux.HandleFunc("GET /export/csv", h.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", h.ExportPDF)
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /ready", h.Health)
	mux.HandleFunc("GET /live", h.Health)
}
