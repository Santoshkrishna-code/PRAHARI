package http

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /documents", h.CreateDocument)
	mux.HandleFunc("GET /documents", h.ListDocuments)
	mux.HandleFunc("GET /documents/{id}", h.GetDocument)
	mux.HandleFunc("POST /documents/{id}/approve", h.ApproveDocument)
	mux.HandleFunc("POST /documents/{id}/publish", h.PublishDocument)
	mux.HandleFunc("POST /documents/{id}/checkout", h.CheckoutDocument)
	mux.HandleFunc("POST /documents/{id}/checkin", h.CheckinDocument)
	mux.HandleFunc("POST /documents/search", h.SearchDocuments)
	mux.HandleFunc("GET /reports", h.GetExecutiveReport)
	mux.HandleFunc("GET /export/csv", h.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", h.ExportPDF)
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /ready", h.Health)
	mux.HandleFunc("GET /live", h.Health)
}
