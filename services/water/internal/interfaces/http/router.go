package http

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /water", h.CreateWaterProfile)
	mux.HandleFunc("GET /water", h.ListWaterProfiles)
	mux.HandleFunc("GET /water/{id}", h.GetWaterProfile)
	mux.HandleFunc("POST /water/meter-reading", h.RecordMeterReading)
	mux.HandleFunc("POST /water/recycling", h.RegisterRecycling)
	mux.HandleFunc("POST /water/forecast", h.ForecastWater)
	mux.HandleFunc("POST /water/optimization", h.CreateOptimization)
	mux.HandleFunc("POST /water/search", h.SearchProfiles)
	mux.HandleFunc("GET /reports", h.GetExecutiveReport)
	mux.HandleFunc("GET /export/csv", h.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", h.ExportPDF)
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /ready", h.Health)
	mux.HandleFunc("GET /live", h.Health)
}
