package http

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /calibrations", h.CreateCalibration)
	mux.HandleFunc("GET /calibrations", h.ListCalibrations)
	mux.HandleFunc("GET /calibrations/{id}", h.GetCalibration)
	mux.HandleFunc("POST /calibrations/{id}/schedule", h.ScheduleCalibration)
	mux.HandleFunc("POST /calibrations/{id}/execute", h.ExecuteCalibration)
	mux.HandleFunc("POST /calibrations/{id}/measurements", h.RecordMeasurements)
	mux.HandleFunc("POST /calibrations/{id}/approve", h.ApproveCalibration)
	mux.HandleFunc("POST /calibrations/search", h.SearchCalibrations)
	mux.HandleFunc("GET /reports", h.GetExecutiveReport)
	mux.HandleFunc("GET /export/csv", h.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", h.ExportPDF)
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /ready", h.Health)
	mux.HandleFunc("GET /live", h.Health)
}
