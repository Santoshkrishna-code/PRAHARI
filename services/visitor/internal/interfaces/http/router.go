package http

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /visitors", h.CreateVisitor)
	mux.HandleFunc("GET /visitors", h.ListVisitors)
	mux.HandleFunc("GET /visitors/{id}", h.GetVisitor)
	mux.HandleFunc("POST /visitors/{id}/approve", h.ApproveVisitor)
	mux.HandleFunc("POST /visitors/{id}/checkin", h.CheckInVisitor)
	mux.HandleFunc("POST /visitors/{id}/checkout", h.CheckOutVisitor)
	mux.HandleFunc("POST /visitors/{id}/muster", h.EvacuateVisitor)
	mux.HandleFunc("POST /visitors/search", h.SearchVisitors)
	mux.HandleFunc("GET /reports", h.GetExecutiveReport)
	mux.HandleFunc("GET /export/csv", h.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", h.ExportPDF)
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /ready", h.Health)
	mux.HandleFunc("GET /live", h.Health)
}
