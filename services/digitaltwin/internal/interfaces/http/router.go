package http

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /twins", h.CreateTwin)
	mux.HandleFunc("GET /twins", h.ListTwins)
	mux.HandleFunc("GET /twins/{id}", h.GetTwin)
	mux.HandleFunc("POST /simulations", h.RunSimulation)
	mux.HandleFunc("POST /playback", h.StartPlayback)
	mux.HandleFunc("GET /overlays", h.GetOverlays)
	mux.HandleFunc("POST /digitaltwin/search", h.SearchTwin)
	mux.HandleFunc("GET /reports", h.GetPerceptionReport)
	mux.HandleFunc("GET /export/csv", h.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", h.ExportPDF)
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /ready", h.Health)
	mux.HandleFunc("GET /live", h.Health)
}
