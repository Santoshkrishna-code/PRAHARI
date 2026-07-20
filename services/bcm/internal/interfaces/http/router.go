package http

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /bcm", h.CreateBCM)
	mux.HandleFunc("GET /bcm", h.ListBCMs)
	mux.HandleFunc("GET /bcm/{id}", h.GetBCM)
	mux.HandleFunc("POST /bcm/{id}/bia", h.ExecuteBIA)
	mux.HandleFunc("POST /bcm/{id}/exercise", h.ScheduleExercise)
	mux.HandleFunc("POST /bcm/{id}/activate", h.ActivatePlan)
	mux.HandleFunc("POST /bcm/{id}/recovery", h.CompleteRecovery)
	mux.HandleFunc("POST /bcm/search", h.SearchBCMs)
	mux.HandleFunc("GET /reports", h.GetExecutiveReport)
	mux.HandleFunc("GET /export/csv", h.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", h.ExportPDF)
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /ready", h.Health)
	mux.HandleFunc("GET /live", h.Health)
}
