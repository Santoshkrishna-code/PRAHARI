package http

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /cameras", h.RegisterCamera)
	mux.HandleFunc("GET /cameras", h.ListCameras)
	mux.HandleFunc("POST /models", h.RegisterModel)
	mux.HandleFunc("POST /inference", h.RunInference)
	mux.HandleFunc("GET /detections", h.GetDetections)
	mux.HandleFunc("POST /alerts", h.CreateAlert)
	mux.HandleFunc("POST /vision/search", h.SearchVisionDetections)
	mux.HandleFunc("GET /reports", h.GetPerceptionReport)
	mux.HandleFunc("GET /export/csv", h.ExportCSV)
	mux.HandleFunc("GET /export/pdf/{id}", h.ExportPDF)
	mux.HandleFunc("GET /health", h.Health)
	mux.HandleFunc("GET /ready", h.Health)
	mux.HandleFunc("GET /live", h.Health)
}
